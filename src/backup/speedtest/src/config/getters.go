package config


func (c ConfigurationConfig) Production() string{
        return  c.production
    }
    func (c ConfigurationConfig) Development() string{
        return  c.development
    }
    func (c ConfigurationConfig) Testing() string{
        return  c.testing
    }
    
    
func (e EmailConfig) Enabled() bool{
        return  e.enabled
    }
    func (e EmailConfig) Host() string{
        return  e.host
    }
    func (e EmailConfig) Username() string{
        return  e.username
    }
    func (e EmailConfig) Password() string{
        return  e.password
    }
    func (e EmailConfig) InsecureSkipVerify() bool{
        return  e.insecureSkipVerify
    }
    func (e EmailConfig) RecipientUsername() string{
        return  e.recipientUsername
    }
    func (e EmailConfig) RecipientPassword() string{
        return  e.recipientPassword
    }
    func (e EmailConfig) RecipientHost() string{
        return  e.recipientHost
    }
    func (e EmailConfig) Subject() string{
        return  e.subject
    }
    func (e EmailConfig) LogMessages() bool{
        return  e.logMessages
    }
    
    
func (h HttpServerConfig) Enabled() bool{
        return  h.enabled
    }
    func (h HttpServerConfig) ServerName() string{
        return  h.serverName
    }
    func (h HttpServerConfig) Bind() string{
        return  h.bind
    }
    func (h HttpServerConfig) Port() int{
        return h.port
    } 
    func (h HttpServerConfig) LogRequests() bool{
        return  h.logRequests
    }
    func (h HttpServerConfig) Cookie() bool{
        return  h.cookie
    }
    
    
func (i IsDevConfig) IsDev() bool{
        return i.isDev
    }
    

    func (j JaegerClientsConfig) Enabled() bool{
        return j.enabled
    }
    func (j JaegerClientsConfig) JaegerServiceName() string{
        return j.jaegerServiceName
    }
    func (j JaegerClientsConfig) JaegerEndpoint() string{
        return j.jaegerEndpoint
    }
    func (j JaegerClientsConfig) LogMessages() bool{
        return j.logMessages
    }
    

    func (m MonetDBClientsConfig) Enabled() bool{
        return m.enabled
    }
    func (m MonetDBClientsConfig) DatabaseHost() string{
        return m.databaseHost
    }
    func (m MonetDBClientsConfig) DatabasePort() int{
        return m.databasePort
    } 
    func (m MonetDBClientsConfig) DatabaseUsername() string{
        return m.databaseUsername
    }
    func (m MonetDBClientsConfig) DatabasePassword() string{
        return m.databasePassword
    }
    func (m MonetDBClientsConfig) DatabaseTimeout() int{
        return m.databaseTimeout
    } 
    func (m MonetDBClientsConfig) DatabaseDriver() string{
        return m.databaseDriver
    }
    func (m MonetDBClientsConfig) DatabaseName() string{
        return m.databaseName
    }
    func (m MonetDBClientsConfig) DatabaseSSLMode() string{
        return m.databaseSSLMode
    }
    func (m MonetDBClientsConfig) MigrationDir() string{
        return m.migrationDir
    }
    func (m MonetDBClientsConfig) DownMigration() bool{
        return m.downMigration
    }
    
func (p PasetoTokenConfig) Enabled() bool{
        return  p.enabled
    }
    func (p PasetoTokenConfig) TokenSymmetricKey() string{
        return  p.tokenSymmetricKey
    }
    func (p PasetoTokenConfig) AccessTokenDuration() int{
        return p.accessTokenDuration
    } 
    func (p PasetoTokenConfig) RefreshTokenDuration() int{
        return p.refreshTokenDuration
    } 
    func (p PasetoTokenConfig) AccessTokenDurationDev() int{
        return p.accessTokenDurationDev
    } 
    func (p PasetoTokenConfig) RefreshTokenDurationDev() int{
        return p.refreshTokenDurationDev
    } 
    func (p PasetoTokenConfig) IsDev() bool{
        return  p.isDev
    }
    
    

    func (p PostgresDBClientsConfig) Enabled() bool{
        return p.enabled
    }
    func (p PostgresDBClientsConfig) DatabaseHost() string{
        return p.databaseHost
    }
    func (p PostgresDBClientsConfig) DatabasePort() int{
        return p.databasePort
    } 
    func (p PostgresDBClientsConfig) DatabaseUsername() string{
        return p.databaseUsername
    }
    func (p PostgresDBClientsConfig) DatabasePassword() string{
        return p.databasePassword
    }
    func (p PostgresDBClientsConfig) DatabaseDriver() string{
        return p.databaseDriver
    }
    func (p PostgresDBClientsConfig) DatabaseName() string{
        return p.databaseName
    }
    func (p PostgresDBClientsConfig) DatabaseTimeout() int{
        return p.databaseTimeout
    } 
    func (p PostgresDBClientsConfig) DatabaseSSLMode() string{
        return p.databaseSSLMode
    }
    func (p PostgresDBClientsConfig) MigrationDir() string{
        return p.migrationDir
    }
    func (p PostgresDBClientsConfig) DownMigration() bool{
        return p.downMigration
    }
    
func (q QueueConfig) Enabled() bool{
        return  q.enabled
    }
    func (q QueueConfig) MaxWorkers() int{
        return q.maxWorkers
    } 
    func (q QueueConfig) LogMessages() bool{
        return  q.logMessages
    }
    
    

    func (r RedisDBClientsConfig) Enabled() bool{
        return r.enabled
    }
    func (r RedisDBClientsConfig) Name() string{
        return r.name
    }
    func (r RedisDBClientsConfig) RedisPassword() string{
        return r.redisPassword
    }
    func (r RedisDBClientsConfig) RedisAddress() string{
        return r.redisAddress
    }
    func (r RedisDBClientsConfig) DatabaseUsername() string{
        return r.databaseUsername
    }
    func (r RedisDBClientsConfig) DatabaseTimeout() int{
        return r.databaseTimeout
    } 
    func (r RedisDBClientsConfig) RedisDatabase() int{
        return r.redisDatabase
    } 
    func (r RedisDBClientsConfig) LogMessages() bool{
        return r.logMessages
    }
    

    func (s ScyllaDBClientsConfig) Enabled() bool{
        return s.enabled
    }
    func (s ScyllaDBClientsConfig) Name() string{
        return s.name
    }
    func (s ScyllaDBClientsConfig) ScyllaHosts() []string{
        return s.scyllaHosts
    } 
    func (s ScyllaDBClientsConfig) Username() string{
        return s.username
    }
    func (s ScyllaDBClientsConfig) Password() string{
        return s.password
    }
    func (s ScyllaDBClientsConfig) DatabaseTimeout() int{
        return s.databaseTimeout
    } 
    func (s ScyllaDBClientsConfig) Keyspace() string{
        return s.keyspace
    }
    func (s ScyllaDBClientsConfig) Class() string{
        return s.class
    }
    func (s ScyllaDBClientsConfig) ReplicationFactor() int{
        return s.replicationFactor
    } 
    func (s ScyllaDBClientsConfig) MigrationDir() string{
        return s.migrationDir
    }
    func (s ScyllaDBClientsConfig) DurableWrites() bool{
        return s.durableWrites
    }
    
func (v VersionConfig)  Version() int{
        return v.version
   }
    

