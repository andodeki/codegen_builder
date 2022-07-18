package config


func (c Config) MarshalYAML() (interface{}, error){
    return configRead{

        Configuration: func() *configurationConfigRead {
            if !c.Configuration.() {
                    return nil
                }
                r := c.Configuration.convertToRead()
                return &r
        }(),
        Email: func() *emailConfigRead {
            if !c.Email.Enabled() {
                    return nil
                }
                r := c.Email.convertToRead()
                return &r
        }(),
        HttpServer: func() *httpServerConfigRead {
            if !c.HttpServer.Enabled() {
                    return nil
                }
                r := c.HttpServer.convertToRead()
                return &r
        }(),
        IsDev: &c.IsDev,
        JaegerClients: func() jaegerClientsConfigReadMap {
            jaegerClients:= make(jaegerClientsConfigReadMap, len(c.JaegerClients))
                for _, c := range c.JaegerClients {
                    if !c.Enabled() {
                        return nil
                    }
                    jaegerClients[c.JaegerServiceName()] = c.convertToRead()
                }
            return jaegerClients
        }(),   
        MonetDBClients: func() monetDBClientsConfigReadMap {
            monetDBClients:= make(monetDBClientsConfigReadMap, len(c.MonetDBClients))
                for _, c := range c.MonetDBClients {
                    if !c.Enabled() {
                        return nil
                    }
                    monetDBClients[c.DatabaseName()] = c.convertToRead()
                }
            return monetDBClients
        }(),   
        PasetoToken: func() *pasetoTokenConfigRead {
            if !c.PasetoToken.Enabled() {
                    return nil
                }
                r := c.PasetoToken.convertToRead()
                return &r
        }(),
        PostgresDBClients: func() postgresDBClientsConfigReadMap {
            postgresDBClients:= make(postgresDBClientsConfigReadMap, len(c.PostgresDBClients))
                for _, c := range c.PostgresDBClients {
                    if !c.Enabled() {
                        return nil
                    }
                    postgresDBClients[c.DatabaseName()] = c.convertToRead()
                }
            return postgresDBClients
        }(),   
        Queue: func() *queueConfigRead {
            if !c.Queue.Enabled() {
                    return nil
                }
                r := c.Queue.convertToRead()
                return &r
        }(),
        RedisDBClients: func() redisDBClientsConfigReadMap {
            redisDBClients:= make(redisDBClientsConfigReadMap, len(c.RedisDBClients))
                for _, c := range c.RedisDBClients {
                    if !c.Enabled() {
                        return nil
                    }
                    redisDBClients[c.Name()] = c.convertToRead()
                }
            return redisDBClients
        }(),   
        ScyllaDBClients: func() scyllaDBClientsConfigReadMap {
            scyllaDBClients:= make(scyllaDBClientsConfigReadMap, len(c.ScyllaDBClients))
                for _, c := range c.ScyllaDBClients {
                    if !c.Enabled() {
                        return nil
                    }
                    scyllaDBClients[c.Name()] = c.convertToRead()
                }
            return scyllaDBClients
        }(),   
        Version: &c.Version, 
        
    }, nil
}

func (c ConfigurationConfig) convertToRead() configurationConfigRead {
    return configurationConfigRead{
        Production: &c.production, 
        Development: &c.development, 
        Testing: &c.testing, 
        
        
    }
}
func (e EmailConfig) convertToRead() emailConfigRead {
    return emailConfigRead{
        Enabled: &e.enabled, 
        Host: &e.host, 
        Username: &e.username, 
        Password: &e.password, 
        InsecureSkipVerify: &e.insecureSkipVerify, 
        RecipientUsername: &e.recipientUsername, 
        RecipientPassword: &e.recipientPassword, 
        RecipientHost: &e.recipientHost, 
        Subject: &e.subject, 
        LogMessages: &e.logMessages, 
        
        
    }
}
func (h HttpServerConfig) convertToRead() httpServerConfigRead {
    return httpServerConfigRead{
        Enabled: &h.enabled, 
        ServerName: &h.serverName, 
        Bind: &h.bind, 
        Port: h.port,
        LogRequests: &h.logRequests, 
        Cookie: &h.cookie, 
        
        
    }
}
func (i IsDevConfig) convertToRead() isDevConfigRead {
    return isDevConfigRead{
        IsDev: i.isDev, //bool
    }
}
func (j JaegerClientsConfig) convertToRead() jaegerClientsConfigRead {
    return jaegerClientsConfigRead{
        
        Enabled: &j.enabled, 
        JaegerServiceName: &j.jaegerServiceName, 
        JaegerEndpoint: &j.jaegerEndpoint, 
        LogMessages: &j.logMessages, 
    }
}
func (m MonetDBClientsConfig) convertToRead() monetDBClientsConfigRead {
    return monetDBClientsConfigRead{
        
        Enabled: &m.enabled, 
        DatabaseHost: &m.databaseHost, 
        DatabasePort: m.databasePort,
        DatabaseUsername: &m.databaseUsername, 
        DatabasePassword: &m.databasePassword, 
        DatabaseTimeout: m.databaseTimeout,
        DatabaseDriver: &m.databaseDriver, 
        DatabaseName: &m.databaseName, 
        DatabaseSSLMode: &m.databaseSSLMode, 
        MigrationDir: &m.migrationDir, 
        DownMigration: &m.downMigration, 
    }
}
func (p PasetoTokenConfig) convertToRead() pasetoTokenConfigRead {
    return pasetoTokenConfigRead{
        Enabled: &p.enabled, 
        TokenSymmetricKey: &p.tokenSymmetricKey, 
        AccessTokenDuration: p.accessTokenDuration,
        RefreshTokenDuration: p.refreshTokenDuration,
        AccessTokenDurationDev: p.accessTokenDurationDev,
        RefreshTokenDurationDev: p.refreshTokenDurationDev,
        IsDev: &p.isDev, 
        
        
    }
}
func (p PostgresDBClientsConfig) convertToRead() postgresDBClientsConfigRead {
    return postgresDBClientsConfigRead{
        
        Enabled: &p.enabled, 
        DatabaseHost: &p.databaseHost, 
        DatabasePort: p.databasePort,
        DatabaseUsername: &p.databaseUsername, 
        DatabasePassword: &p.databasePassword, 
        DatabaseDriver: &p.databaseDriver, 
        DatabaseName: &p.databaseName, 
        DatabaseTimeout: p.databaseTimeout,
        DatabaseSSLMode: &p.databaseSSLMode, 
        MigrationDir: &p.migrationDir, 
        DownMigration: &p.downMigration, 
    }
}
func (q QueueConfig) convertToRead() queueConfigRead {
    return queueConfigRead{
        Enabled: &q.enabled, 
        MaxWorkers: q.maxWorkers,
        LogMessages: &q.logMessages, 
        
        
    }
}
func (r RedisDBClientsConfig) convertToRead() redisDBClientsConfigRead {
    return redisDBClientsConfigRead{
        
        Enabled: &r.enabled, 
        Name: &r.name, 
        RedisPassword: &r.redisPassword, 
        RedisAddress: &r.redisAddress, 
        DatabaseUsername: &r.databaseUsername, 
        DatabaseTimeout: r.databaseTimeout,
        RedisDatabase: r.redisDatabase,
        LogMessages: &r.logMessages, 
    }
}
func (s ScyllaDBClientsConfig) convertToRead() scyllaDBClientsConfigRead {
    return scyllaDBClientsConfigRead{
        
        Enabled: &s.enabled, 
        Name: &s.name, 
        ScyllaHosts: s.scyllaHosts,
        Username: &s.username, 
        Password: &s.password, 
        DatabaseTimeout: s.databaseTimeout,
        Keyspace: &s.keyspace, 
        Class: &s.class, 
        ReplicationFactor: s.replicationFactor,
        MigrationDir: &s.migrationDir, 
        DurableWrites: &s.durableWrites, 
    }
}
func (v VersionConfig) convertToRead() versionConfigRead {
    return versionConfigRead{
        Version: v.version, //int 
        
    }
}

