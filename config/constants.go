package config

const (
	// For Redis
	REDIS_HOST = "redis.yce:32379"
	// REDIS_PORT = "32379"
	// REDIS_SERVER = "redis.yce:32379"
	MAX_IDLE     = 1
	MAX_ACTIVE   = 10
	IDEL_TIMEOUT = 180

	// For MySQL
	MAX_POOL_SIZE        = 20
	DATABASE_DRIVER      = "mysql"
	DB_HOST              = "mysql.yce:3306"
	DB_USER              = "root"
	DB_PASSWORD          = "root"
	DB_NAME              = "yce"
	DB_CONNECTION_SUFFIX = "?parseTime=true"
	DELAY_MILLISECONDS   = 5000

	// For Registry
	REGISTRY_HOST = "img.reg.3g"
	REGISTRY_PORT = "15000"
	REGISTRY_CERT = "domain.crt"

	QA = "QA"
	SIAGENT_HOST = "10.151.30.227"
	SIAGENT_PORT = "8081"
)
