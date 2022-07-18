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
    _ "github.com/MonetDB/MonetDB-Go/src"
)





type MonetDBClients struct {
    logger *util.Logger
    config *config.MonetDBClients
    client *sqlx.DB
}

type MonetDBClientsInterface interface{
    MonetDBClient()  *sqlx.DB
    Health(ctx context.Context)  error
    
    Transaction(ctx context.Context, fn func(ctx context.Context) error)  error
    migrateDb(ctx context.Context, logger *util.Logger, config *config.MonetDBClients)  error
    waitForDB(ctx context.Context)  error 
}

func (m *MonetDBClients) MonetDBClient() *sqlx.DB{
    //DBClient
    return m.client
    
    
}  


func (m *MonetDBClients) Health(ctx context.Context) error{
    //Health
    delay := time.NewTicker(100 * time.Millisecond)
	timeoutExceeded := time.After(300 * time.Millisecond)
    for {
		select {
		case <-timeoutExceeded:
            return util.NewError("database failed: timeoutExceeded")
        case <-delay.C:
            err := func() error {
                if int32(m.client.Stats().OpenConnections) < int32(1) {
					return util.NewError("open connections size below minimum")
				}
                //Check if database is running
                if err := m.waitForDB(ctx); err != nil {
					return errors.Wrap(err, "datasources.Ping: could not ping to database")
				}
                if err := ctx.Err(); err != nil {
					return err
				}
                var ok bool
                _ = ok
                //m
                if err := m.client.GetContext(ctx, &ok, `SELECT true;`); !ok || err != nil {
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


func RunMonetDB(ctx context.Context, retries int, logger *util.Logger, config *config.MonetDBClients) (client MonetDBClientsInterface,err error){
    //Run
    logger.Info("connecting to monet database!")
    delay := time.NewTicker(1 * time.Second)
	timeoutExceeded := time.After((time.Duration(retries) * time.Second))
    //DatabaseSSLMode()
    dsn := url.URL{
		User:   url.UserPassword(config.DatabaseUsername(), config.DatabasePassword()),
		Host:   fmt.Sprintf("%s:%s", config.DatabaseHost(), config.DatabasePort()),
		Path:   config.DatabaseName(),
    }
    q := dsn.Query()
	dsn.RawQuery = q.Encode()
	dbURL := dsn.String()

    
    for {
		select {
		case <-timeoutExceeded:
            return &MonetDBClients{}, util.NewError("database failed: timeoutExceeded")
        case <-delay.C:
			logger.Info("trying to connect to the database")
            conn := sqlx.MustConnect(config.DatabaseDriver(), trimFirstRune(dbURL))
            conn.SetMaxOpenConns(32)
			logger.Info("connected to the monet database")

            client = &MonetDBClients{logger: logger, config: config, client: conn}
            return client, nil
            }
    }

    
}  


func (m *MonetDBClients) Transaction(ctx context.Context, fn func(ctx context.Context) error) error{
    //Transaction
    return nil
    
}  


func (m *MonetDBClients) migrateDb(ctx context.Context, logger *util.Logger, config *config.MonetDBClients) error{
    //migrateDb
    return nil
    
}  


func (m *MonetDBClients) waitForDB(ctx context.Context) error{
    //waitForDB
    ready := make(chan struct{})
	go func() {
		for {
			if err := m.client.Ping(); err == nil {
				close(ready)
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()
	select {
	case <-ready:
		return nil
	case <-time.After(time.Duration(*m.config.DatabaseTimeout()) * time.Millisecond):
		return errors.New("database not ready")
	}
    

    
}  



func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i+1:]
}


