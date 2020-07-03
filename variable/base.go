package variable

// The place that you can share the various variables accross the whole project.

import (
	"os"
	"fmt"
)

var (
	Env =  os.Getenv("ENV")
	dbName     = os.Getenv("DB_NAME")
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbPort     = os.Getenv("DB_PORT")
	dbHost     = os.Getenv("DB_HOST")
	sslMode    = os.Getenv("SSL_MODE")
	maxPoolConns = os.Getenv("MAX_POOL_CONNECTIONS")
	minPoolConns = os.Getenv("MIN_POOL_CONNECTIONS")
	DBURL = buildDBURL()
	Concurrency = os.Getenv("GOROUTINE_CONCURRENCY")
	AMQPURL = os.Getenv("CLOUDAMQP_URL")
	TestDBURL  = os.Getenv("TEST_DB_DSN")
	FixturesURI = os.Getenv("FIXTURES_URI")
)

func buildDBURL() (dbURL string) {
	switch Env {
	case "development":
		dbURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&pool_max_conns=%s&pool_min_conns=%s", dbUser, dbPassword, dbHost, dbPort, dbName, sslMode, maxPoolConns, minPoolConns)
	default:
		dbURL = fmt.Sprintf("%s?sslmode=require&pool_max_conns=%s&pool_min_conns=%s", os.Getenv("DATABASE_URL"), maxPoolConns, minPoolConns)
	}
	return dbURL
}
