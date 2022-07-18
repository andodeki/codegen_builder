package config



import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"runtime"
	"sort"
	"strings"

	"github.com/namsral/flag"
	"github.com/pkg/errors"

	"gopkg.in/yaml.v3"
)
func ReadConfigFile(exe, source string) (config Config, err []error) {
	yamlStr, e := ioutil.ReadFile(source)
	if e != nil {
		return config, []error{fmt.Errorf("cannot read configuration: %v; use see `%s --help`", err, exe)}
	}

	return ReadConfig(yamlStr)
}

func ReadConfig(yamlStr []byte) (config Config, err []error) {
	var configRead configRead

	e := yaml.Unmarshal(yamlStr, &configRead)
	if e != nil {
		return config, []error{errors.Wrap(e, "cannot parse yaml: %")}
	}

	return configRead.TransformAndValidate()
}

func (c Config) PrintConfig() (err error) {
	newYamlStr, err := yaml.Marshal(c)
	if err != nil {
		return errors.Wrap(err, "cannot encode yaml again: %v")
	}
	log.Print("config: use the following config:")
	for _, line := range strings.Split(string(newYamlStr), "\n") {
		log.Print("config: ", line)
	}
	return nil
}

/*func (c configRead) TransformAndValidate() (ret Config, err []error) {
	var e []error

}
*/


func (c configRead) TransformAndValidate() (ret Config, err []error) {
	var e []error
    
        ret.Configuration, e = c.Configuration.TransformAndValidate()
        err = append(err, e...)
        
        
        ret.Email, e = c.Email.TransformAndValidate()
        err = append(err, e...)
        
        
        ret.HttpServer, e = c.HttpServer.TransformAndValidate()
        err = append(err, e...)
        
        
        if c.IsDev != nil && *c.IsDev {
            ret.IsDev = *c.IsDev//bool
        }
        ret.JaegerClients, e = c.JaegerClients.TransformAndValidate() 
        err = append(err, e...)//[]*JaegerClientsConfig 
        
        ret.MonetDBClients, e = c.MonetDBClients.TransformAndValidate() 
        err = append(err, e...)//[]*MonetDBClientsConfig 
        
        ret.PasetoToken, e = c.PasetoToken.TransformAndValidate()
        err = append(err, e...)
        
        
        ret.PostgresDBClients, e = c.PostgresDBClients.TransformAndValidate() 
        err = append(err, e...)//[]*PostgresDBClientsConfig 
        
        ret.Queue, e = c.Queue.TransformAndValidate()
        err = append(err, e...)
        
        
        ret.RedisDBClients, e = c.RedisDBClients.TransformAndValidate() 
        err = append(err, e...)//[]*RedisDBClientsConfig 
        
        ret.ScyllaDBClients, e = c.ScyllaDBClients.TransformAndValidate() 
        err = append(err, e...)//[]*ScyllaDBClientsConfig 
        
        if c.Version == nil {
            ret.Version = *c.Version//int
        } else {
            ret.Version = *c.Version
            if ret.Version != 0 {
                err = append(err, fmt.Errorf("version=%d is not supported", ret.Version))
        }
        
    }
    return
}


func (c *configurationConfigRead) TransformAndValidate() (ret ConfigurationConfig, err []error) {
    if c == nil {
        return
    }
    if *c.Enabled {
        if len(ret.production) < 1 {
            ret.production = *c.Production
        }
        if len(ret.production) < 1 {
            err = append(err, fmt.Errorf("ProductionConfig->%s->Production must not be empty", c.Production))
        }
        if len(ret.development) < 1 {
            ret.development = *c.Development
        }
        if len(ret.development) < 1 {
            err = append(err, fmt.Errorf("DevelopmentConfig->%s->Development must not be empty", c.Production))
        }
        if len(ret.testing) < 1 {
            ret.testing = *c.Testing
        }
        if len(ret.testing) < 1 {
            err = append(err, fmt.Errorf("TestingConfig->%s->Testing must not be empty", c.Production))
        }
        }
	return
    //Production: &c.production, 
}

func (e *emailConfigRead) TransformAndValidate() (ret EmailConfig, err []error) {
    if e == nil {
        return
    }
    if *e.Enabled {
        if e.Enabled != nil && *e.Enabled {
            ret.enabled = *e.Enabled
        }
                if len(ret.host) < 1 {
            ret.host = *e.Host
        }
        if len(ret.host) < 1 {
            err = append(err, fmt.Errorf("HostConfig->%s->Host must not be empty", e.Enabled))
        }
        if len(ret.username) < 1 {
            ret.username = *e.Username
        }
        if len(ret.username) < 1 {
            err = append(err, fmt.Errorf("UsernameConfig->%s->Username must not be empty", e.Enabled))
        }
        if len(ret.password) < 1 {
            ret.password = *e.Password
        }
        if len(ret.password) < 1 {
            err = append(err, fmt.Errorf("PasswordConfig->%s->Password must not be empty", e.Enabled))
        }
        if e.InsecureSkipVerify != nil && *e.InsecureSkipVerify {
            ret.insecureSkipVerify = *e.InsecureSkipVerify
        }
                if len(ret.recipientUsername) < 1 {
            ret.recipientUsername = *e.RecipientUsername
        }
        if len(ret.recipientUsername) < 1 {
            err = append(err, fmt.Errorf("RecipientUsernameConfig->%s->RecipientUsername must not be empty", e.Enabled))
        }
        if len(ret.recipientPassword) < 1 {
            ret.recipientPassword = *e.RecipientPassword
        }
        if len(ret.recipientPassword) < 1 {
            err = append(err, fmt.Errorf("RecipientPasswordConfig->%s->RecipientPassword must not be empty", e.Enabled))
        }
        if len(ret.recipientHost) < 1 {
            ret.recipientHost = *e.RecipientHost
        }
        if len(ret.recipientHost) < 1 {
            err = append(err, fmt.Errorf("RecipientHostConfig->%s->RecipientHost must not be empty", e.Enabled))
        }
        if len(ret.subject) < 1 {
            ret.subject = *e.Subject
        }
        if len(ret.subject) < 1 {
            err = append(err, fmt.Errorf("SubjectConfig->%s->Subject must not be empty", e.Enabled))
        }
        if e.LogMessages != nil && *e.LogMessages {
            ret.logMessages = *e.LogMessages
        }
                }
	return
    //Enabled: &e.enabled, 
}

func (h *httpServerConfigRead) TransformAndValidate() (ret HttpServerConfig, err []error) {
    if h == nil {
        return
    }
    if *h.Enabled {
        if h.Enabled != nil && *h.Enabled {
            ret.enabled = *h.Enabled
        }
                if len(ret.serverName) < 1 {
            ret.serverName = *h.ServerName
        }
        if len(ret.serverName) < 1 {
            err = append(err, fmt.Errorf("ServerNameConfig->%s->ServerName must not be empty", h.Enabled))
        }
        if len(ret.bind) < 1 {
            ret.bind = *h.Bind
        }
        if len(ret.bind) < 1 {
            err = append(err, fmt.Errorf("BindConfig->%s->Bind must not be empty", h.Enabled))
        }
        if ret.port < 1 {
            ret.port = h.Port
        }
        if ret.port < 1 {
            err = append(err, fmt.Errorf("PortConfig->%s->Port must not be empty", h.Enabled))
        }
            if h.LogRequests != nil && *h.LogRequests {
            ret.logRequests = *h.LogRequests
        }
                if h.Cookie != nil && *h.Cookie {
            ret.cookie = *h.Cookie
        }
                }
	return
    //Enabled: &h.enabled, 
}

func (j jaegerClientsConfigReadMap) TransformAndValidate() (ret []*JaegerClientsConfig, err []error) {
    //if len(j) < 1 {
		//return ret, []error{fmt.Errorf("JaegerClients section must no be empty")}
	//}

	ret = make([]*JaegerClientsConfig, len(j))
	jj := 0
	for _, name := range j.getOrderedKeys() {
		r, e := j[name].TransformAndValidate(name)
		ret[jj] = &r
		err = append(err, e...)
		jj++
	}
	return
}
func (j jaegerClientsConfigRead) TransformAndValidate(name string) (ret JaegerClientsConfig, err []error) {
	ret.enabled = *j.Enabled
	if *j.Enabled {
    //Enabled: &j.enabled, 
        if j.Enabled != nil && *j.Enabled {
			ret.enabled = *j.Enabled
		}
        
        if len(ret.jaegerServiceName) < 1 {
	 	    ret.jaegerServiceName = *j.JaegerServiceName
	    }
        if len(ret.jaegerServiceName) < 1 {
			err = append(err, fmt.Errorf("JaegerServiceNameConfig->%s->JaegerServiceName must not be empty", j.JaegerServiceName))
		}
        
        if len(ret.jaegerEndpoint) < 1 {
	 	    ret.jaegerEndpoint = *j.JaegerEndpoint
	    }
        if len(ret.jaegerEndpoint) < 1 {
			err = append(err, fmt.Errorf("JaegerEndpointConfig->%s->JaegerEndpoint must not be empty", j.JaegerEndpoint))
		}
        
        if j.LogMessages != nil && *j.LogMessages {
			ret.logMessages = *j.LogMessages
		}
        
    }

	return
}
func (j jaegerClientsConfigReadMap) getOrderedKeys() (ret []string) {
	ret = make([]string, len(j))
	i := 0
	for k := range j {
		ret[i] = k
		i++
	}
	sort.Strings(ret)
	return
}


func (m monetDBClientsConfigReadMap) TransformAndValidate() (ret []*MonetDBClientsConfig, err []error) {
    //if len(m) < 1 {
		//return ret, []error{fmt.Errorf("MonetDBClients section must no be empty")}
	//}

	ret = make([]*MonetDBClientsConfig, len(m))
	jj := 0
	for _, name := range m.getOrderedKeys() {
		r, e := m[name].TransformAndValidate(name)
		ret[jj] = &r
		err = append(err, e...)
		jj++
	}
	return
}
func (m monetDBClientsConfigRead) TransformAndValidate(name string) (ret MonetDBClientsConfig, err []error) {
	ret.enabled = *m.Enabled
	if *m.Enabled {
    //Enabled: &m.enabled, 
        if m.Enabled != nil && *m.Enabled {
			ret.enabled = *m.Enabled
		}
        
        if len(ret.databaseHost) < 1 {
	 	    ret.databaseHost = *m.DatabaseHost
	    }
        if len(ret.databaseHost) < 1 {
			err = append(err, fmt.Errorf("DatabaseHostConfig->%s->DatabaseHost must not be empty", m.DatabaseHost))
		}
        if ret.databasePort < 1 {
			ret.databasePort = m.DatabasePort
		}
        if ret.databasePort < 1 {
			err = append(err, fmt.Errorf("DatabasePortConfig->%s->DatabasePort must not be empty", m.DatabasePort))
		}
        
        if len(ret.databaseUsername) < 1 {
	 	    ret.databaseUsername = *m.DatabaseUsername
	    }
        if len(ret.databaseUsername) < 1 {
			err = append(err, fmt.Errorf("DatabaseUsernameConfig->%s->DatabaseUsername must not be empty", m.DatabaseUsername))
		}
        
        if len(ret.databasePassword) < 1 {
	 	    ret.databasePassword = *m.DatabasePassword
	    }
        if len(ret.databasePassword) < 1 {
			err = append(err, fmt.Errorf("DatabasePasswordConfig->%s->DatabasePassword must not be empty", m.DatabasePassword))
		}
        if ret.databaseTimeout < 1 {
			ret.databaseTimeout = m.DatabaseTimeout
		}
        if ret.databaseTimeout < 1 {
			err = append(err, fmt.Errorf("DatabaseTimeoutConfig->%s->DatabaseTimeout must not be empty", m.DatabaseTimeout))
		}
        
        if len(ret.databaseDriver) < 1 {
	 	    ret.databaseDriver = *m.DatabaseDriver
	    }
        if len(ret.databaseDriver) < 1 {
			err = append(err, fmt.Errorf("DatabaseDriverConfig->%s->DatabaseDriver must not be empty", m.DatabaseDriver))
		}
        
        if len(ret.databaseName) < 1 {
	 	    ret.databaseName = *m.DatabaseName
	    }
        if len(ret.databaseName) < 1 {
			err = append(err, fmt.Errorf("DatabaseNameConfig->%s->DatabaseName must not be empty", m.DatabaseName))
		}
        
        if len(ret.databaseSSLMode) < 1 {
	 	    ret.databaseSSLMode = *m.DatabaseSSLMode
	    }
        if len(ret.databaseSSLMode) < 1 {
			err = append(err, fmt.Errorf("DatabaseSSLModeConfig->%s->DatabaseSSLMode must not be empty", m.DatabaseSSLMode))
		}
        
        if len(ret.migrationDir) < 1 {
	 	    ret.migrationDir = *m.MigrationDir
	    }
        if len(ret.migrationDir) < 1 {
			err = append(err, fmt.Errorf("MigrationDirConfig->%s->MigrationDir must not be empty", m.MigrationDir))
		}
        
        if m.DownMigration != nil && *m.DownMigration {
			ret.downMigration = *m.DownMigration
		}
        
    }

	return
}
func (m monetDBClientsConfigReadMap) getOrderedKeys() (ret []string) {
	ret = make([]string, len(m))
	i := 0
	for k := range m {
		ret[i] = k
		i++
	}
	sort.Strings(ret)
	return
}



func (p *pasetoTokenConfigRead) TransformAndValidate() (ret PasetoTokenConfig, err []error) {
    if p == nil {
        return
    }
    if *p.Enabled {
        if p.Enabled != nil && *p.Enabled {
            ret.enabled = *p.Enabled
        }
                if len(ret.tokenSymmetricKey) < 1 {
            ret.tokenSymmetricKey = *p.TokenSymmetricKey
        }
        if len(ret.tokenSymmetricKey) < 1 {
            err = append(err, fmt.Errorf("TokenSymmetricKeyConfig->%s->TokenSymmetricKey must not be empty", p.Enabled))
        }
        if ret.accessTokenDuration < 1 {
            ret.accessTokenDuration = p.AccessTokenDuration
        }
        if ret.accessTokenDuration < 1 {
            err = append(err, fmt.Errorf("AccessTokenDurationConfig->%s->AccessTokenDuration must not be empty", p.Enabled))
        }
            if ret.refreshTokenDuration < 1 {
            ret.refreshTokenDuration = p.RefreshTokenDuration
        }
        if ret.refreshTokenDuration < 1 {
            err = append(err, fmt.Errorf("RefreshTokenDurationConfig->%s->RefreshTokenDuration must not be empty", p.Enabled))
        }
            if ret.accessTokenDurationDev < 1 {
            ret.accessTokenDurationDev = p.AccessTokenDurationDev
        }
        if ret.accessTokenDurationDev < 1 {
            err = append(err, fmt.Errorf("AccessTokenDurationDevConfig->%s->AccessTokenDurationDev must not be empty", p.Enabled))
        }
            if ret.refreshTokenDurationDev < 1 {
            ret.refreshTokenDurationDev = p.RefreshTokenDurationDev
        }
        if ret.refreshTokenDurationDev < 1 {
            err = append(err, fmt.Errorf("RefreshTokenDurationDevConfig->%s->RefreshTokenDurationDev must not be empty", p.Enabled))
        }
            if p.IsDev != nil && *p.IsDev {
            ret.isDev = *p.IsDev
        }
                }
	return
    //Enabled: &p.enabled, 
}
func (p postgresDBClientsConfigReadMap) TransformAndValidate() (ret []*PostgresDBClientsConfig, err []error) {
    //if len(p) < 1 {
		//return ret, []error{fmt.Errorf("PostgresDBClients section must no be empty")}
	//}

	ret = make([]*PostgresDBClientsConfig, len(p))
	jj := 0
	for _, name := range p.getOrderedKeys() {
		r, e := p[name].TransformAndValidate(name)
		ret[jj] = &r
		err = append(err, e...)
		jj++
	}
	return
}
func (p postgresDBClientsConfigRead) TransformAndValidate(name string) (ret PostgresDBClientsConfig, err []error) {
	ret.enabled = *p.Enabled
	if *p.Enabled {
    //Enabled: &p.enabled, 
        if p.Enabled != nil && *p.Enabled {
			ret.enabled = *p.Enabled
		}
        
        if len(ret.databaseHost) < 1 {
	 	    ret.databaseHost = *p.DatabaseHost
	    }
        if len(ret.databaseHost) < 1 {
			err = append(err, fmt.Errorf("DatabaseHostConfig->%s->DatabaseHost must not be empty", p.DatabaseHost))
		}
        if ret.databasePort < 1 {
			ret.databasePort = p.DatabasePort
		}
        if ret.databasePort < 1 {
			err = append(err, fmt.Errorf("DatabasePortConfig->%s->DatabasePort must not be empty", p.DatabasePort))
		}
        
        if len(ret.databaseUsername) < 1 {
	 	    ret.databaseUsername = *p.DatabaseUsername
	    }
        if len(ret.databaseUsername) < 1 {
			err = append(err, fmt.Errorf("DatabaseUsernameConfig->%s->DatabaseUsername must not be empty", p.DatabaseUsername))
		}
        
        if len(ret.databasePassword) < 1 {
	 	    ret.databasePassword = *p.DatabasePassword
	    }
        if len(ret.databasePassword) < 1 {
			err = append(err, fmt.Errorf("DatabasePasswordConfig->%s->DatabasePassword must not be empty", p.DatabasePassword))
		}
        
        if len(ret.databaseDriver) < 1 {
	 	    ret.databaseDriver = *p.DatabaseDriver
	    }
        if len(ret.databaseDriver) < 1 {
			err = append(err, fmt.Errorf("DatabaseDriverConfig->%s->DatabaseDriver must not be empty", p.DatabaseDriver))
		}
        
        if len(ret.databaseName) < 1 {
	 	    ret.databaseName = *p.DatabaseName
	    }
        if len(ret.databaseName) < 1 {
			err = append(err, fmt.Errorf("DatabaseNameConfig->%s->DatabaseName must not be empty", p.DatabaseName))
		}
        if ret.databaseTimeout < 1 {
			ret.databaseTimeout = p.DatabaseTimeout
		}
        if ret.databaseTimeout < 1 {
			err = append(err, fmt.Errorf("DatabaseTimeoutConfig->%s->DatabaseTimeout must not be empty", p.DatabaseTimeout))
		}
        
        if len(ret.databaseSSLMode) < 1 {
	 	    ret.databaseSSLMode = *p.DatabaseSSLMode
	    }
        if len(ret.databaseSSLMode) < 1 {
			err = append(err, fmt.Errorf("DatabaseSSLModeConfig->%s->DatabaseSSLMode must not be empty", p.DatabaseSSLMode))
		}
        
        if len(ret.migrationDir) < 1 {
	 	    ret.migrationDir = *p.MigrationDir
	    }
        if len(ret.migrationDir) < 1 {
			err = append(err, fmt.Errorf("MigrationDirConfig->%s->MigrationDir must not be empty", p.MigrationDir))
		}
        
        if p.DownMigration != nil && *p.DownMigration {
			ret.downMigration = *p.DownMigration
		}
        
    }

	return
}
func (p postgresDBClientsConfigReadMap) getOrderedKeys() (ret []string) {
	ret = make([]string, len(p))
	i := 0
	for k := range p {
		ret[i] = k
		i++
	}
	sort.Strings(ret)
	return
}



func (q *queueConfigRead) TransformAndValidate() (ret QueueConfig, err []error) {
    if q == nil {
        return
    }
    if *q.Enabled {
        if q.Enabled != nil && *q.Enabled {
            ret.enabled = *q.Enabled
        }
                if ret.maxWorkers < 1 {
            ret.maxWorkers = q.MaxWorkers
        }
        if ret.maxWorkers < 1 {
            err = append(err, fmt.Errorf("MaxWorkersConfig->%s->MaxWorkers must not be empty", q.Enabled))
        }
            if q.LogMessages != nil && *q.LogMessages {
            ret.logMessages = *q.LogMessages
        }
                }
	return
    //Enabled: &q.enabled, 
}
func (r redisDBClientsConfigReadMap) TransformAndValidate() (ret []*RedisDBClientsConfig, err []error) {
    //if len(r) < 1 {
		//return ret, []error{fmt.Errorf("RedisDBClients section must no be empty")}
	//}

	ret = make([]*RedisDBClientsConfig, len(r))
	jj := 0
	for _, name := range r.getOrderedKeys() {
		r, e := r[name].TransformAndValidate(name)
		ret[jj] = &r
		err = append(err, e...)
		jj++
	}
	return
}
func (r redisDBClientsConfigRead) TransformAndValidate(name string) (ret RedisDBClientsConfig, err []error) {
	ret.enabled = *r.Enabled
	if *r.Enabled {
    //Enabled: &r.enabled, 
        if r.Enabled != nil && *r.Enabled {
			ret.enabled = *r.Enabled
		}
        
        if len(ret.name) < 1 {
	 	    ret.name = *r.Name
	    }
        if len(ret.name) < 1 {
			err = append(err, fmt.Errorf("NameConfig->%s->Name must not be empty", r.Name))
		}
        
        if len(ret.redisPassword) < 1 {
	 	    ret.redisPassword = *r.RedisPassword
	    }
        if len(ret.redisPassword) < 1 {
			err = append(err, fmt.Errorf("RedisPasswordConfig->%s->RedisPassword must not be empty", r.RedisPassword))
		}
        
        if len(ret.redisAddress) < 1 {
	 	    ret.redisAddress = *r.RedisAddress
	    }
        if len(ret.redisAddress) < 1 {
			err = append(err, fmt.Errorf("RedisAddressConfig->%s->RedisAddress must not be empty", r.RedisAddress))
		}
        
        if len(ret.databaseUsername) < 1 {
	 	    ret.databaseUsername = *r.DatabaseUsername
	    }
        if len(ret.databaseUsername) < 1 {
			err = append(err, fmt.Errorf("DatabaseUsernameConfig->%s->DatabaseUsername must not be empty", r.DatabaseUsername))
		}
        if ret.databaseTimeout < 1 {
			ret.databaseTimeout = r.DatabaseTimeout
		}
        if ret.databaseTimeout < 1 {
			err = append(err, fmt.Errorf("DatabaseTimeoutConfig->%s->DatabaseTimeout must not be empty", r.DatabaseTimeout))
		}
        if ret.redisDatabase < 1 {
			ret.redisDatabase = r.RedisDatabase
		}
        if ret.redisDatabase < 1 {
			err = append(err, fmt.Errorf("RedisDatabaseConfig->%s->RedisDatabase must not be empty", r.RedisDatabase))
		}
        
        if r.LogMessages != nil && *r.LogMessages {
			ret.logMessages = *r.LogMessages
		}
        
    }

	return
}
func (r redisDBClientsConfigReadMap) getOrderedKeys() (ret []string) {
	ret = make([]string, len(r))
	i := 0
	for k := range r {
		ret[i] = k
		i++
	}
	sort.Strings(ret)
	return
}


func (s scyllaDBClientsConfigReadMap) TransformAndValidate() (ret []*ScyllaDBClientsConfig, err []error) {
    //if len(s) < 1 {
		//return ret, []error{fmt.Errorf("ScyllaDBClients section must no be empty")}
	//}

	ret = make([]*ScyllaDBClientsConfig, len(s))
	jj := 0
	for _, name := range s.getOrderedKeys() {
		r, e := s[name].TransformAndValidate(name)
		ret[jj] = &r
		err = append(err, e...)
		jj++
	}
	return
}
func (s scyllaDBClientsConfigRead) TransformAndValidate(name string) (ret ScyllaDBClientsConfig, err []error) {
	ret.enabled = *s.Enabled
	if *s.Enabled {
    //Enabled: &s.enabled, 
        if s.Enabled != nil && *s.Enabled {
			ret.enabled = *s.Enabled
		}
        
        if len(ret.name) < 1 {
	 	    ret.name = *s.Name
	    }
        if len(ret.name) < 1 {
			err = append(err, fmt.Errorf("NameConfig->%s->Name must not be empty", s.Name))
		}
        if len(ret.scyllaHosts) < 1 {
	 	    ret.scyllaHosts = s.ScyllaHosts
	    }
        if len(ret.scyllaHosts) < 1 {
			err = append(err, fmt.Errorf("ScyllaHostsConfig->%s->ScyllaHosts must not be empty", s.ScyllaHosts))
		}
        
        if len(ret.username) < 1 {
	 	    ret.username = *s.Username
	    }
        if len(ret.username) < 1 {
			err = append(err, fmt.Errorf("UsernameConfig->%s->Username must not be empty", s.Username))
		}
        
        if len(ret.password) < 1 {
	 	    ret.password = *s.Password
	    }
        if len(ret.password) < 1 {
			err = append(err, fmt.Errorf("PasswordConfig->%s->Password must not be empty", s.Password))
		}
        if ret.databaseTimeout < 1 {
			ret.databaseTimeout = s.DatabaseTimeout
		}
        if ret.databaseTimeout < 1 {
			err = append(err, fmt.Errorf("DatabaseTimeoutConfig->%s->DatabaseTimeout must not be empty", s.DatabaseTimeout))
		}
        
        if len(ret.keyspace) < 1 {
	 	    ret.keyspace = *s.Keyspace
	    }
        if len(ret.keyspace) < 1 {
			err = append(err, fmt.Errorf("KeyspaceConfig->%s->Keyspace must not be empty", s.Keyspace))
		}
        
        if len(ret.class) < 1 {
	 	    ret.class = *s.Class
	    }
        if len(ret.class) < 1 {
			err = append(err, fmt.Errorf("ClassConfig->%s->Class must not be empty", s.Class))
		}
        if ret.replicationFactor < 1 {
			ret.replicationFactor = s.ReplicationFactor
		}
        if ret.replicationFactor < 1 {
			err = append(err, fmt.Errorf("ReplicationFactorConfig->%s->ReplicationFactor must not be empty", s.ReplicationFactor))
		}
        
        if len(ret.migrationDir) < 1 {
	 	    ret.migrationDir = *s.MigrationDir
	    }
        if len(ret.migrationDir) < 1 {
			err = append(err, fmt.Errorf("MigrationDirConfig->%s->MigrationDir must not be empty", s.MigrationDir))
		}
        
        if s.DurableWrites != nil && *s.DurableWrites {
			ret.durableWrites = *s.DurableWrites
		}
        
    }

	return
}
func (s scyllaDBClientsConfigReadMap) getOrderedKeys() (ret []string) {
	ret = make([]string, len(s))
	i := 0
	for k := range s {
		ret[i] = k
		i++
	}
	sort.Strings(ret)
	return
}




