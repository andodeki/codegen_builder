package datasources


import (
    "time"
    "fmt"
    "context"
	"unicode/utf8"
	"net/url"
	"database/sql"

    "github.com/andodeki/propertylisting/src/util"
    "github.com/andodeki/propertylisting/src/config"


    "github.com/pkg/errors"
    "embed"

	"github.com/jmoiron/sqlx"
    "github.com/sirupsen/logrus"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

var db *sql.DB
type ContextKey string
const CONTEXT_TRANSACTION_KEY ContextKey = "database:tx"

//go:embed migrations/*.sql
var fs embed.FS




type PostgresDBClients struct {
    logger *util.Logger
    config *config.PostgresDBClients
    client *sqlx.DB
}

type PostgresDBClientsInterface interface{
    PostgresDBClient()  *sqlx.DB
    Health(ctx context.Context)  error
    
    Transaction(ctx context.Context, fn func(ctx context.Context) error)  error
    migrateDb(ctx context.Context, logger *util.Logger, config *config.PostgresDBClients)  error
    waitForDB(ctx context.Context)  error 
}

func (p *PostgresDBClients) PostgresDBClient() *sqlx.DB{
    //DBClient
    return p.client
    
    
}  


func (p *PostgresDBClients) Health(ctx context.Context) error{
    //Health
    delay := time.NewTicker(100 * time.Millisecond)
	timeoutExceeded := time.After(300 * time.Millisecond)
    for {
		select {
		case <-timeoutExceeded:
            return util.NewError("database failed: timeoutExceeded")
        case <-delay.C:
            err := func() error {
                if int32(p.client.Stats().OpenConnections) < int32(1) {
					return util.NewError("open connections size below minimum")
				}
                //Check if database is running
                if err := p.waitForDB(ctx); err != nil {
					return errors.Wrap(err, "datasources.Ping: could not ping to database")
				}
                if err := ctx.Err(); err != nil {
					return err
				}
                var ok bool
                _ = ok
                //p
                if err := p.client.GetContext(ctx, &ok, `SELECT true;`); !ok || err != nil {
					return errors.Wrap(err, "datasources.Health: SELECT true;")
				}
                if err := ctx.Err(); err != nil {
					return err
				}
                return nil
            }()
            if err != nil {
				return util.NewError("failed: delay case").Wrap(err)
			}
			return nil
        }
    }
    
}  


func RunPostgresDB(ctx context.Context, retries int, logger *util.Logger, config *config.PostgresDBClients) (client PostgresDBClientsInterface,err error){
    //Run
    logger.Info("connecting to postgres database!")
    delay := time.NewTicker(1 * time.Second)
	timeoutExceeded := time.After((time.Duration(retries) * time.Second))
    //DatabaseSSLMode()
    dsn := url.URL{
        Scheme: config.DatabaseDriver(),
		User:   url.UserPassword(config.DatabaseUsername(), config.DatabasePassword()),
		Host:   fmt.Sprintf("%s:%s", config.DatabaseHost(), config.DatabasePort()),
		Path:   config.DatabaseName(),
    }
    q := dsn.Query()
    q.Add("sslmode", config.DatabaseSSLMode())
	dsn.RawQuery = q.Encode()
	dbURL := dsn.String()

    
    for {
		select {
		case <-timeoutExceeded:
            return &PostgresDBClients{}, util.NewError("database failed: timeoutExceeded")
        case <-delay.C:
			logger.Info("trying to connect to the database")
            conn := sqlx.MustConnect(config.DatabaseDriver(), dbURL)
            conn.SetMaxOpenConns(32)
			logger.Info("connected to the postgres database")

            
            //Check if database is running
			//if err := client.waitForDB(ctx); err != nil {
			//	return nil, errors.Wrap(err, "datasources.Ping: could not ping to database")
			//}

			//Do database migration
			if err := client.migrateDb(ctx, logger, config); err != nil {
				return nil, errors.Wrap(err, "datasources.Migrate: could not migrate database")
			}
            client = &PostgresDBClients{logger: logger, config: config, client: conn}
            return client, nil
            }
    }

    
}  


func (p *PostgresDBClients) Transaction(ctx context.Context, fn func(ctx context.Context) error) error{
    //Transaction
    if ctx.Value(CONTEXT_TRANSACTION_KEY) != nil {
		err := fn(ctx)
		if err != nil {
			return util.NewError("database transaction failed").WrapWithDepth(1, err)
		}
		return nil
	}

	transaction, err := p.client.BeginTx(ctx, nil)
	if err != nil {
		return util.NewError("database transaction failed").WrapWithDepth(1, err)
	}

	if err := ctx.Err(); err != nil {
		return util.NewError("database transaction failed").WrapWithDepth(1, err)
	}

	defer func() {
		if pan := recover(); pan != nil {
			_ = transaction.Rollback()
			panic(pan)
		}
	}()

	err = fn(context.WithValue(ctx, CONTEXT_TRANSACTION_KEY, transaction))
	if err != nil {
		errR := transaction.Rollback()
		if errR == nil {
			return util.NewError("database transaction failed").AsWithDepth(1, err)
		}

		err := errors.CombineErrors(err, errR)
		return util.NewError("database transaction failed").WrapWithDepth(1, err)
	}

	if err := ctx.Err(); err != nil {
		errR := transaction.Rollback()
		err := errors.CombineErrors(err, errR)
		return util.NewError("database transaction failed").WrapWithDepth(1, err)
	}

	if err := transaction.Commit(); err != nil {
		errR := transaction.Rollback()
		err := errors.CombineErrors(err, errR)
		return util.NewError("database transaction failed").WrapWithDepth(1, err)
	}

	if err := ctx.Err(); err != nil {
		errR := transaction.Rollback()
		err := errors.CombineErrors(err, errR)
		return util.NewError("database transaction failed").WrapWithDepth(1, err)
	}
    return nil
    
    return nil
    
}  


func (p *PostgresDBClients) migrateDb(ctx context.Context, logger *util.Logger, config *config.PostgresDBClients) error{
    //migrateDb
    sourceInstance, err := iofs.New(fs, "migrations")
	if err != nil {
		return err
	}
    
    db = p.client

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return errors.Wrap(err, "connecting to database")
	}
    migrator, err := migrate.NewWithInstance("iofs", sourceInstance, "postgres", driver)
	if err != nil {
		return errors.Wrap(err, "creating migrator")
	}
    if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
		return errors.Wrap(err, "executing up migration")
	}

	version, dirty, err := migrator.Version()
	if err != nil {
		return errors.Wrap(err, "getting migration version")
	}
	logrus.WithFields(logrus.Fields{
		"version": version,
		"dirty":   dirty,
	}).Debug("database migrated")
    return sourceInstance.Close()
    
    
}  


func (p *PostgresDBClients) waitForDB(ctx context.Context) error{
    //waitForDB
    ready := make(chan struct{})
	go func() {
		for {
			if err := p.client.Ping(); err == nil {
				close(ready)
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()
	select {
	case <-ready:
		return nil
	case <-time.After(time.Duration(*p.config.DatabaseTimeout()) * time.Millisecond):
		return errors.New("database not ready")
	}
    

    
}  




