package config


type Config struct{

    Configuration ConfigurationConfig `yaml:"Configuration"`
    Email EmailConfig `yaml:"Email"`
    HttpServer HttpServerConfig `yaml:"HttpServer"`
    IsDev bool `yaml:"IsDev"`
    JaegerClients  []*JaegerClientsConfig `yaml:"JaegerClients"`
    MonetDBClients  []*MonetDBClientsConfig `yaml:"MonetDBClients"`
    PasetoToken PasetoTokenConfig `yaml:"PasetoToken"`
    PostgresDBClients  []*PostgresDBClientsConfig `yaml:"PostgresDBClients"`
    Queue QueueConfig `yaml:"Queue"`
    RedisDBClients  []*RedisDBClientsConfig `yaml:"RedisDBClients"`
    ScyllaDBClients  []*ScyllaDBClientsConfig `yaml:"ScyllaDBClients"`
    Version int `yaml:"Version"`
    
}

type ConfigurationConfig struct{
     production string
     development string
     testing string
    
    
    
}
type EmailConfig struct{
     enabled bool
     host string
     username string
     password string
     insecureSkipVerify bool
     recipientUsername string
     recipientPassword string
     recipientHost string
     subject string
     logMessages bool
    
    
    
}
type HttpServerConfig struct{
     enabled bool
     serverName string
     bind string
     port int 
     logRequests bool
     cookie bool
    
    
    
}
type IsDevConfig struct{
    isDev bool
}
type JaegerClientsConfig struct{
    
    enabled bool
    jaegerServiceName string
    jaegerEndpoint string
    logMessages bool
    
}
type MonetDBClientsConfig struct{
    
    enabled bool
    databaseHost string
    databasePort int 
    databaseUsername string
    databasePassword string
    databaseTimeout int 
    databaseDriver string
    databaseName string
    databaseSSLMode string
    migrationDir string
    downMigration bool
    
}
type PasetoTokenConfig struct{
     enabled bool
     tokenSymmetricKey string
     accessTokenDuration int 
     refreshTokenDuration int 
     accessTokenDurationDev int 
     refreshTokenDurationDev int 
     isDev bool
    
    
    
}
type PostgresDBClientsConfig struct{
    
    enabled bool
    databaseHost string
    databasePort int 
    databaseUsername string
    databasePassword string
    databaseDriver string
    databaseName string
    databaseTimeout int 
    databaseSSLMode string
    migrationDir string
    downMigration bool
    
}
type QueueConfig struct{
     enabled bool
     maxWorkers int 
     logMessages bool
    
    
    
}
type RedisDBClientsConfig struct{
    
    enabled bool
    name string
    redisPassword string
    redisAddress string
    databaseUsername string
    databaseTimeout int 
    redisDatabase int 
    logMessages bool
    
}
type ScyllaDBClientsConfig struct{
    
    enabled bool
    name string
    scyllaHosts []string 
    username string
    password string
    databaseTimeout int 
    keyspace string
    class string
    replicationFactor int 
    migrationDir string
    durableWrites bool
    
}
type VersionConfig struct{
    version int 
    
}






type configRead struct{

    Configuration *configurationConfigRead `yaml:"Configuration"`
    Email *emailConfigRead `yaml:"Email"`
    HttpServer *httpServerConfigRead `yaml:"HttpServer"`
    IsDev *bool `yaml:"IsDev"`
    JaegerClients  jaegerClientsConfigReadMap `yaml:"JaegerClients"`
    MonetDBClients  monetDBClientsConfigReadMap `yaml:"MonetDBClients"`
    PasetoToken *pasetoTokenConfigRead `yaml:"PasetoToken"`
    PostgresDBClients  postgresDBClientsConfigReadMap `yaml:"PostgresDBClients"`
    Queue *queueConfigRead `yaml:"Queue"`
    RedisDBClients  redisDBClientsConfigReadMap `yaml:"RedisDBClients"`
    ScyllaDBClients  scyllaDBClientsConfigReadMap `yaml:"ScyllaDBClients"`
    Version *int `yaml:"Version"`
    
}

type configurationConfigRead struct{
     Production *string `yaml:"Production"`
     Development *string `yaml:"Development"`
     Testing *string `yaml:"Testing"`
    
    
    
}
type emailConfigRead struct{
     Enabled *bool `yaml:"Enabled"`
     Host *string `yaml:"Host"`
     Username *string `yaml:"Username"`
     Password *string `yaml:"Password"`
     InsecureSkipVerify *bool `yaml:"InsecureSkipVerify"`
     RecipientUsername *string `yaml:"RecipientUsername"`
     RecipientPassword *string `yaml:"RecipientPassword"`
     RecipientHost *string `yaml:"RecipientHost"`
     Subject *string `yaml:"Subject"`
     LogMessages *bool `yaml:"LogMessages"`
    
    
    
}
type httpServerConfigRead struct{
     Enabled *bool `yaml:"Enabled"`
     ServerName *string `yaml:"ServerName"`
     Bind *string `yaml:"Bind"`
     Port int  `yaml:"Port"`
     LogRequests *bool `yaml:"LogRequests"`
     Cookie *bool `yaml:"Cookie"`
    
    
    
}
type isDevConfigRead struct{
    IsDev bool`yaml:"IsDev"`
    
}
type jaegerClientsConfigRead struct{
    
    Enabled *bool `yaml:"Enabled"`
    JaegerServiceName *string `yaml:"JaegerServiceName"`
    JaegerEndpoint *string `yaml:"JaegerEndpoint"`
    LogMessages *bool `yaml:"LogMessages"`
    
}
type monetDBClientsConfigRead struct{
    
    Enabled *bool `yaml:"Enabled"`
    DatabaseHost *string `yaml:"DatabaseHost"`
    DatabasePort int  `yaml:"DatabasePort"`
    DatabaseUsername *string `yaml:"DatabaseUsername"`
    DatabasePassword *string `yaml:"DatabasePassword"`
    DatabaseTimeout int  `yaml:"DatabaseTimeout"`
    DatabaseDriver *string `yaml:"DatabaseDriver"`
    DatabaseName *string `yaml:"DatabaseName"`
    DatabaseSSLMode *string `yaml:"DatabaseSSLMode"`
    MigrationDir *string `yaml:"MigrationDir"`
    DownMigration *bool `yaml:"DownMigration"`
    
}
type pasetoTokenConfigRead struct{
     Enabled *bool `yaml:"Enabled"`
     TokenSymmetricKey *string `yaml:"TokenSymmetricKey"`
     AccessTokenDuration int  `yaml:"AccessTokenDuration"`
     RefreshTokenDuration int  `yaml:"RefreshTokenDuration"`
     AccessTokenDurationDev int  `yaml:"AccessTokenDurationDev"`
     RefreshTokenDurationDev int  `yaml:"RefreshTokenDurationDev"`
     IsDev *bool `yaml:"IsDev"`
    
    
    
}
type postgresDBClientsConfigRead struct{
    
    Enabled *bool `yaml:"Enabled"`
    DatabaseHost *string `yaml:"DatabaseHost"`
    DatabasePort int  `yaml:"DatabasePort"`
    DatabaseUsername *string `yaml:"DatabaseUsername"`
    DatabasePassword *string `yaml:"DatabasePassword"`
    DatabaseDriver *string `yaml:"DatabaseDriver"`
    DatabaseName *string `yaml:"DatabaseName"`
    DatabaseTimeout int  `yaml:"DatabaseTimeout"`
    DatabaseSSLMode *string `yaml:"DatabaseSSLMode"`
    MigrationDir *string `yaml:"MigrationDir"`
    DownMigration *bool `yaml:"DownMigration"`
    
}
type queueConfigRead struct{
     Enabled *bool `yaml:"Enabled"`
     MaxWorkers int  `yaml:"MaxWorkers"`
     LogMessages *bool `yaml:"LogMessages"`
    
    
    
}
type redisDBClientsConfigRead struct{
    
    Enabled *bool `yaml:"Enabled"`
    Name *string `yaml:"Name"`
    RedisPassword *string `yaml:"RedisPassword"`
    RedisAddress *string `yaml:"RedisAddress"`
    DatabaseUsername *string `yaml:"DatabaseUsername"`
    DatabaseTimeout int  `yaml:"DatabaseTimeout"`
    RedisDatabase int  `yaml:"RedisDatabase"`
    LogMessages *bool `yaml:"LogMessages"`
    
}
type scyllaDBClientsConfigRead struct{
    
    Enabled *bool `yaml:"Enabled"`
    Name *string `yaml:"Name"`
    ScyllaHosts []string  `yaml:"ScyllaHosts"`
    Username *string `yaml:"Username"`
    Password *string `yaml:"Password"`
    DatabaseTimeout int  `yaml:"DatabaseTimeout"`
    Keyspace *string `yaml:"Keyspace"`
    Class *string `yaml:"Class"`
    ReplicationFactor int  `yaml:"ReplicationFactor"`
    MigrationDir *string `yaml:"MigrationDir"`
    DurableWrites *bool `yaml:"DurableWrites"`
    
}
type versionConfigRead struct{
    Version int `yaml:"Version"`
    
}



type jaegerClientsConfigReadMap map[string]jaegerClientsConfigRead
type monetDBClientsConfigReadMap map[string]monetDBClientsConfigRead
type postgresDBClientsConfigReadMap map[string]postgresDBClientsConfigRead
type redisDBClientsConfigReadMap map[string]redisDBClientsConfigRead
type scyllaDBClientsConfigReadMap map[string]scyllaDBClientsConfigRead


