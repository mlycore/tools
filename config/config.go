package config

import (
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	// MySQL
	ENV_DB_HOST = "DB_HOST"
	ENV_DB_USER = "DB_USER"
	ENV_DB_PASS = "DB_PASS"
	ENV_DB_NAME = "DB_NAME"

	// Redis
	ENV_REDIS_HOST = "REDIS_HOST"
	// ENV_REDIS_PORT = "REDIS_PORT"
	ENV_REDIS_MAX_IDLE_CONN   = "REDIS_MAX_IDLE_CONN"
	ENV_REDIS_MAX_ACTIVE_CONN = "REDIS_MAX_ACTIVE_CONN"
	ENV_REDIS_IDLE_TIMEOUT    = "REDIS_IDLE_TIMEOUT"

	// Registry
	ENV_REGISTRY_HOST = "REGISTRY_HOST"
	ENV_REGISTRY_PORT = "REGISTRY_PORT"
	ENV_REGISTRY_CERT = "REGISTRY_CERT"

	ENV_QA = "QA"
	ENV_SIAGENTHOST = "SIAGENTHOST"
	ENV_SIAGENTPORT = "SIAGNETPORT"
)

var once sync.Once

var instance *Config

func Instance() *Config {
	once.Do(func() {
		instance = new(Config)
	})
	return instance
}

type Config struct {
	DbHost string `json:"DbHost"`
	DbUser string `json:"DbUser"`
	DbPass string `json:"DbPass"`
	DbName string `json:"DbName"`

	RedisHost string `json:"redisHost"`
	// RedisPort string `json:"redisPort"`
	RedisMaxIdleConn   int           `json:"redisMaxIdleConn"`
	RedisMaxActiveConn int           `json:"redisMaxActiveConf"`
	RedisIdleTimeout   time.Duration `json:"redisIdleTimeout"`

	RegistryHost string `json:"RegistryHost"`
	RegistryPort string `json:"RegistryPort"`
	RegistryCert string `json:"RegistryCert"`

	QA           string `json:"QA"`
	SiAgentHost  string `json:"SiAgentHost"`
	SiAgentPort  string `json:"SiAgentPort"`
}

func (conf *Config) Load() {

	// MySQL
	if conf.DbHost = os.Getenv(ENV_DB_HOST); strings.EqualFold(conf.DbHost, "") {
		conf.DbHost = DB_HOST
	}

	if conf.DbUser = os.Getenv(ENV_DB_USER); strings.EqualFold(conf.DbUser, "") {
		conf.DbUser = DB_USER
	}

	if conf.DbPass = os.Getenv(ENV_DB_PASS); strings.EqualFold(conf.DbPass, "") {
		conf.DbPass = DB_PASSWORD
	}

	if conf.DbName = os.Getenv(ENV_DB_NAME); strings.EqualFold(conf.DbName, "") {
		conf.DbName = DB_NAME
	}

	// Redis
	if conf.RedisHost = os.Getenv(ENV_REDIS_HOST); strings.EqualFold(conf.RedisHost, "") {
		conf.RedisHost = REDIS_HOST
	}
	/*
		if conf.RedisPort = os.Getenv(ENV_REDIS_PORT); strings.EqualFold(conf.RedisPort, "") {
			conf.RedisPort = REDIS_PORT
		}
	*/

	if redisMaxIdleConn := os.Getenv(ENV_REDIS_MAX_IDLE_CONN); strings.EqualFold(redisMaxIdleConn, "") {
		conf.RedisMaxIdleConn = MAX_IDLE
	} else {
		rmic, err := strconv.Atoi(redisMaxIdleConn)
		if err != nil {
			conf.RedisMaxIdleConn = MAX_IDLE
		} else {
			conf.RedisMaxIdleConn = rmic
		}
	}

	if redisMaxActiveConn := os.Getenv(ENV_REDIS_MAX_ACTIVE_CONN); strings.EqualFold(redisMaxActiveConn, "") {
		rmac, err := strconv.Atoi(redisMaxActiveConn)
		if err != nil {
			conf.RedisMaxActiveConn = MAX_ACTIVE
		} else {
			conf.RedisMaxActiveConn = rmac
		}
	}

	if redisIdleTimeout := os.Getenv(ENV_REDIS_IDLE_TIMEOUT); strings.EqualFold(redisIdleTimeout, "") {
		rit, err := strconv.Atoi(redisIdleTimeout)
		if err != nil {
			conf.RedisIdleTimeout = time.Duration(IDEL_TIMEOUT) * time.Second
		} else {
			conf.RedisIdleTimeout = time.Duration(rit) * time.Second
		}
	}

	// Registry
	if conf.RegistryHost = os.Getenv(ENV_REGISTRY_HOST); strings.EqualFold(conf.RegistryHost, "") {
		conf.RegistryHost = REGISTRY_HOST
	}

	if conf.RegistryPort = os.Getenv(ENV_REGISTRY_PORT); strings.EqualFold(conf.RegistryPort, "") {
		conf.RegistryPort = REGISTRY_PORT
	}

	if conf.RegistryCert = os.Getenv(ENV_REGISTRY_CERT); strings.EqualFold(conf.RegistryCert, "") {
		conf.RegistryCert = REGISTRY_CERT
	}

	// Check cert file exists
	_, err := os.Open(conf.RegistryCert)
	if err != nil {
		log.Fatalf("Registry Cert file Not found: cert=%s", conf.RegistryCert)
	}

	if conf.QA = os.Getenv(ENV_QA); strings.EqualFold(conf.QA, QA) {
		if conf.SiAgentHost = os.Getenv(ENV_SIAGENTHOST); strings.EqualFold(conf.SiAgentHost, "") {
			conf.SiAgentHost = SIAGENT_HOST
		}
		if conf.SiAgentPort = os.Getenv(ENV_SIAGENTPORT); strings.EqualFold(conf.SiAgentPort, "") {
			conf.SiAgentPort = SIAGENT_PORT
		}
	}

}

func (conf *Config) GetDbEndpoint() string {
	endpoint := conf.DbUser + ":" + conf.DbPass + "@tcp(" + conf.DbHost + ")/" + conf.DbName + DB_CONNECTION_SUFFIX
	log.Infof("MySQL Cluster Endpoint: endpoint=%s", endpoint)
	return endpoint
}

func (conf *Config) GetRedisEndpoint() string {
	log.Infof("Redis Cluster Endpoint: endpoint=%s", conf.RedisHost)
	return conf.RedisHost
}
