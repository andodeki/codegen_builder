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
    "github.com/andodeki/propertylisting/src/datasources/migrations/cqlmigrations"
    "github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
    cqlMigrate "github.com/scylladb/gocqlx/v2/migrate"
)



var configf = struct {
	DB       gocql.ClusterConfig
	Password gocql.PasswordAuthenticator
}{}


type ScyllaDBClients struct {
    logger *util.Logger
    config *config.ScyllaDBClients
    clientSession  *gocql.Session
    clientx *gocqlx.Session
}

type ScyllaDBClientsInterface interface{
    ScyllaDBClient()  *gocqlx.Session
    Health(ctx context.Context)  error
    
    Transaction(ctx context.Context, fn func(ctx context.Context) error)  error
    migrateDb(ctx context.Context, logger *util.Logger, config *config.ScyllaDBClients)  error
    waitForDB(ctx context.Context)  error 
}

func (s *ScyllaDBClients) ScyllaDBClient() *gocqlx.Session{
    //DBClient
    return s.clientx
    
    
}  


func (s *ScyllaDBClients) Health(ctx context.Context) error{
    //Health
    delay := time.NewTicker(100 * time.Millisecond)
	timeoutExceeded := time.After(300 * time.Millisecond)
    for {
		select {
		case <-timeoutExceeded:
            return util.NewError("database failed: timeoutExceeded")
        case <-delay.C:
            err := func() error {if err := ctx.Err(); err != nil {
					return err
				}
                var ok bool
                _ = ok
                //s
                if err := s.clientSession.Query("SELECT true").Scan(&ok); !ok || err != nil {
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


func RunScyllaDB(ctx context.Context, retries int, logger *util.Logger, config *config.ScyllaDBClients) (client ScyllaDBClientsInterface,err error){
    //Run
    logger.Info("connecting to scylla database!")
    delay := time.NewTicker(1 * time.Second)
	timeoutExceeded := time.After((time.Duration(retries) * time.Second))
    //
    for {
		select {
		case <-timeoutExceeded:
            return &ScyllaDBClients{}, util.NewError("database failed: timeoutExceeded")
        case <-delay.C:
			logger.Info("trying to connect to the database")
            
            logger.Info("trying to connect to the database")
            //Do database migration
			if err := client.migrateDb(ctx, logger, config); err != nil {
				return nil, errors.Wrap(err, "datasources.Migrate: could not migrate to scylla database")
			}

            clientSession, err := gocql.NewSession(Config(config))
			if err != nil {
				return nil, errors.Wrap(err, "datasources.Ping: could not create new session")
			}
			cfg := Config(config)
			cfg.Keyspace = config.Keyspace()
			// fmt.Printf("cfg.Keyspace: %v\n", cfg.Keyspace)
			clientSessionX, err := gocqlx.WrapSession(cfg.CreateSession())
			if err != nil {
				return nil, errors.Wrap(err, "datasources.Ping: could not create wrapped session")
			}
            client = &ScyllaDBClients{logger: logger, config: config, clientSession: clientSession, clientx: &clientSessionX}
            return client, nil
			logger.Info("connected to the scylla database")
            }
    }

    
}  


func (s *ScyllaDBClients) Transaction(ctx context.Context, fn func(ctx context.Context) error) error{
    //Transaction
    return nil
    
}  


func (s *ScyllaDBClients) migrateDb(ctx context.Context, logger *util.Logger, config *config.ScyllaDBClients) error{
    //migrateDb
    KeySpaceCQL := fmt.Sprintf(`CREATE KEYSPACE IF NOT EXISTS %s WITH replication = { 'class': '%s', 'replication_factor': '%d' } AND durable_writes = %t;`, 
    config.Keyspace(), 
    config.Class(), 
    config.ReplicationFactor(), 
    config.DurableWrites(),
    )
	//const KeySpaceCQL = "CREATE KEYSPACE IF NOT EXISTS propertylisting WITH replication = { 'class': 'NetworkTopologyStrategy', 'replication_factor': '3' } AND durable_writes = TRUE;"
	// fmt.Printf("KeySpaceCQL: %v\n", KeySpaceCQL)
	// fmt.Printf("config.Keyspace(): %v\n", config.Keyspace())
	mig := config.MigrationDir()

	createKeyspace(KeySpaceCQL, config, logger)
	migrateKeyspace(ctx, config.Keyspace(), config, mig, logger)
	printKeyspaceMetadata(config.Keyspace(), config, logger)

	return nil
    
}  


func (s *ScyllaDBClients) waitForDB(ctx context.Context) error{
    //waitForDB
    return nil
    

    
}  





func Config(config *config.ScyllaDBClients) gocql.ClusterConfig {
	configf.DB = *gocql.NewCluster()

	configf.DB.Consistency = gocql.LocalOne
	configf.DB.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())
	configf.DB.Hosts = config.ScyllaHosts()
	configf.DB.Timeout = 60 * time.Second
	configf.DB.ConnectTimeout = 5 * time.Second
	configf.Password.Username = ""
	configf.Password.Username = ""

	if configf.Password.Username != "" {
		configf.DB.Authenticator = configf.Password
	}
	var t = configf.DB
	if configf.Password.Username != "" {
		t.Authenticator = configf.Password
	}
	return t
}

// Session returns new session
func Session(config *config.ScyllaDBClients) (*gocql.Session, error) {
	return gocql.NewSession(Config(config))
}

// Keyspace returns new session with specified keyspace
func Keyspacef(keySpace string, config *config.ScyllaDBClients) (gocqlx.Session, error) {
	cfg := Config(config)
	cfg.Keyspace = keySpace
	return gocqlx.WrapSession(gocql.NewSession(cfg))
}

func createKeyspace(KeySpaceCQL string, cfg *config.ScyllaDBClients, logger *util.Logger) {
	ses, err := Session(cfg)
	if err != nil {
		logger.Fatalf("session: ", err)
	}
	defer ses.Close()

	if err := ses.Query(KeySpaceCQL).Exec(); err != nil {
		logger.Infof("ensure keyspace exists: %v", err)
	}
}

func migrateKeyspace(ctx context.Context, keySpace string, config *config.ScyllaDBClients, mig string, logger *util.Logger) {
	sesx, err := Keyspacef(keySpace, config)
	if err != nil {
		logger.Fatalf("session: ", err)
	}
	defer sesx.Close()
	if err := cqlMigrate.FromFS(context.Background(), sesx, cqlmigrations.Files); err != nil {
		logger.Infof("migrate: %v", err)
	}
}

func printKeyspaceMetadata(keyspace string, config *config.ScyllaDBClients, logger *util.Logger) {
	sesx, err := Keyspacef(keyspace, config)
	if err != nil {
		logger.Fatalf("session: ", err)
	}

	defer sesx.Close()

	m, err := sesx.KeyspaceMetadata(keyspace)
	if err != nil {
		logger.Infof("keyspace metadata: %v", err)
	}

	logger.Printf("Keyspace metadata = %+v\n", m)
}
