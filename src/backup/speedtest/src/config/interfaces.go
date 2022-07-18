package config


type Configuration interface{
     Production() string
     Development() string
     Testing() string
    
    
}

type Email interface{
     Enabled() bool
     Host() string
     Username() string
     Password() string
     InsecureSkipVerify() bool
     RecipientUsername() string
     RecipientPassword() string
     RecipientHost() string
     Subject() string
     LogMessages() bool
    
    
}

type HttpServer interface{
     Enabled() bool
     ServerName() string
     Bind() string
     Port() int 
     LogRequests() bool
     Cookie() bool
    
    
}

type IsDev interface{
    IsDev() bool
}

type JaegerClients interface{
    
    Enabled() bool
    JaegerServiceName() string
    JaegerEndpoint() string
    LogMessages() bool
    
}

type MonetDBClients interface{
    
    Enabled() bool
    DatabaseHost() string
    DatabasePort() int 
    DatabaseUsername() string
    DatabasePassword() string
    DatabaseTimeout() int 
    DatabaseDriver() string
    DatabaseName() string
    DatabaseSSLMode() string
    MigrationDir() string
    DownMigration() bool
    
}

type PasetoToken interface{
     Enabled() bool
     TokenSymmetricKey() string
     AccessTokenDuration() int 
     RefreshTokenDuration() int 
     AccessTokenDurationDev() int 
     RefreshTokenDurationDev() int 
     IsDev() bool
    
    
}

type PostgresDBClients interface{
    
    Enabled() bool
    DatabaseHost() string
    DatabasePort() int 
    DatabaseUsername() string
    DatabasePassword() string
    DatabaseDriver() string
    DatabaseName() string
    DatabaseTimeout() int 
    DatabaseSSLMode() string
    MigrationDir() string
    DownMigration() bool
    
}

type Queue interface{
     Enabled() bool
     MaxWorkers() int 
     LogMessages() bool
    
    
}

type RedisDBClients interface{
    
    Enabled() bool
    Name() string
    RedisPassword() string
    RedisAddress() string
    DatabaseUsername() string
    DatabaseTimeout() int 
    RedisDatabase() int 
    LogMessages() bool
    
}

type ScyllaDBClients interface{
    
    Enabled() bool
    Name() string
    ScyllaHosts() []string 
    Username() string
    Password() string
    DatabaseTimeout() int 
    Keyspace() string
    Class() string
    ReplicationFactor() int 
    MigrationDir() string
    DurableWrites() bool
    
}

type Version interface{
    Version() int 
    
}

