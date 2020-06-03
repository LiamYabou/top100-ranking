package app

import (
	"os"
	"github.com/LiamYabou/top100-ranking/pkg/logger"
	"github.com/LiamYabou/top100-ranking/pkg/db"
	"github.com/LiamYabou/top100-ranking/pkg/variable"
	"github.com/LiamYabou/top100-ranking/pkg/rabbitmq"
	"github.com/streadway/amqp"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	DBPool *pgxpool.Pool
	AMQPConn *amqp.Connection
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
	DBPool, err = db.Open()
	if err != nil {
		logger.Error("Failed to connect the DB.", err)
	}
	AMQPConn, err = rabbitmq.Open()
	if err != nil {
		logger.Error("Failed to connect the RabbitMQ.", err)
	}
}