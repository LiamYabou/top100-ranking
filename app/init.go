package app

import (
	"os"
	"github.com/LiamYabou/top100-pkg/logger"
	"github.com/LiamYabou/top100-pkg/db"
	"github.com/LiamYabou/top100-ranking/variable"
	"github.com/LiamYabou/top100-pkg/rabbitmq"
	"github.com/LiamYabou/top100-pkg/monitor"
	"github.com/streadway/amqp"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var (
	DBpool *pgxpool.Pool
	AMQPconn *amqp.Connection
	NewRelicMonitor *newrelic.Application
	file     *os.File
)

func init() {
	var err error
	switch variable.Env {
	case "development":
		file, err = logger.SetDevConfigs()
		if err != nil {
			logger.Error("Failed to set the configs of logger.", err)
		}
	case "staging":
		logger.SetStagingConfigs()
	case "production":
		logger.SetProductionConfigs()
	}
	DBpool, err = db.Open(variable.DBURL)
	if err != nil {
		logger.Error("Failed to connect the DB.", err)
	}
	AMQPconn, err = rabbitmq.Open(variable.AMQPURL)
	if err != nil {
		logger.Error("Failed to connect the RabbitMQ.", err)
	}
	NewRelicMonitor, err = monitor.InitNewRelic(variable.Env, variable.AppName, variable.NewRelicLicenseKey)
	if err != nil {
        logger.Error("unable to create New Relic Application.", err)
	}
	err = monitor.InitSentry(variable.Env)
	if err != nil {
		logger.Error("Unable to configure Sentry.", err)
	}
}
