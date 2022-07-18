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
    "runtime"
    "github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)





type RedisDBClients struct {
    logger *util.Logger
    config *config.RedisDBClients
    client   *redis.Client
    Cache  *cache.Cache
}

type RedisDBClientsInterface interface{
    RedisDBClient()  *redis.Client
    Health(ctx context.Context)  error
    
    Transaction(ctx context.Context, fn func(ctx context.Context) error)  error
    migrateDb(ctx context.Context, logger *util.Logger, config *config.RedisDBClients)  error
    waitForDB(ctx context.Context)  error 
}

func (r *RedisDBClients) RedisDBClient() *redis.Client{
    //DBClient
    return r.client
    
    
}  


func (r *RedisDBClients) Health(ctx context.Context) error{
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
                //r
                
                //Check if database is running
                if err := r.waitForDB(ctx); err != nil {
					return errors.Wrap(err, "datasources.Ping: could not ping to database")
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


func RunRedisDB(ctx context.Context, retries int, logger *util.Logger, config *config.RedisDBClients) (client RedisDBClientsInterface,err error){
    //Run
    logger.Info("connecting to redis database!")
    delay := time.NewTicker(1 * time.Second)
	timeoutExceeded := time.After((time.Duration(retries) * time.Second))
    //
	host := config.RedisAddress()
    rdb := &redis.Options{
		Addr:         host,
		Password:     config.RedisPassword(),
		DB:           0,
		MaxRetries:   3,
		DialTimeout:  time.Duration(15) * time.Second, // 15 S
		ReadTimeout:  time.Duration(15) * time.Second, // 15 S
		PoolTimeout:  time.Duration(15) * time.Second, // 15 S
		WriteTimeout: time.Duration(15) * time.Second, // 15 S
		MinIdleConns: 1,
		PoolSize:     10 * runtime.GOMAXPROCS(-1),
	}
    
    for {
		select {
		case <-timeoutExceeded:
            return &RedisDBClients{}, util.NewError("database failed: timeoutExceeded")
        case <-delay.C:
			logger.Info("trying to connect to the database")
            
            logger.Info("trying to connect to the cache")
            pool := redis.NewClient(rdb)

			result, err := pool.Ping(ctx).Result()
			if err == nil && result == "PONG" {
				logger.Info("connected to the cache")

				cache := cache.New(&cache.Options{
					Redis:        pool,
					LocalCache:   cache.NewTinyLFU(1000, time.Minute),
					StatsEnabled: "dev" == "dev",
				})

                client = &RedisDBClients{logger: logger, config: config, client: pool, Cache: cache}
                return client, nil
			}
			logger.Info("connected to the redis database")
            }
    }

    
}  


func (r *RedisDBClients) Transaction(ctx context.Context, fn func(ctx context.Context) error) error{
    //Transaction
    return nil
    
}  


func (r *RedisDBClients) migrateDb(ctx context.Context, logger *util.Logger, config *config.RedisDBClients) error{
    //migrateDb
    return nil
    
}  


func (r *RedisDBClients) waitForDB(ctx context.Context) error{
    //waitForDB
    ready := make(chan struct{})
	go func() {
		for {
			if err := r.client.Ping(); err == nil {
				close(ready)
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()
	select {
	case <-ready:
		return nil
	case <-time.After(time.Duration(*r.config.DatabaseTimeout()) * time.Millisecond):
		return errors.New("database not ready")
	}
    
    return nil
    

    
}  




