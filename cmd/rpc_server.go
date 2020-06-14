package main

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/LiamYabou/top100-ranking/app"
	"github.com/LiamYabou/top100-ranking/preference"
	"github.com/LiamYabou/top100-pkg/logger"
	"github.com/LiamYabou/top100-ranking/api"
	"github.com/streadway/amqp"
)

func main() {
	defer app.Finalize()
	opts := &preference.Options{
		DB: app.DBpool,
		AMQP: app.AMQPconn, 
	}
	opts = preference.LoadOptions(preference.WithOptions(*opts))
	performRPCserver(opts)
}

func performRPCserver(opts *preference.Options) {
	ch, err := opts.AMQP.Channel()
	if err != nil {
		logger.Error("Failed to open a channel.", err)
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"rpc_queue", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		logger.Error("Failed to declare a queue.", err)
	}
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		logger.Error("Failed to set QoS.", err)
	}
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		logger.Error("Failed to register a consumer.", err)
	}
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			args := strings.Split(string(d.Body), "/")
			// args[0] represents the action of the consumer;
			// args[1] represents the id of the category row;
			// args[2] represents the number of the page.
			categoryId, _ := strconv.Atoi(args[1]) 
			page, _ := strconv.Atoi(args[2])
			response := api.FindRankings(categoryId, page, opts)
			err = ch.Publish(
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "application/json",
					CorrelationId: d.CorrelationId,
					Body:          []byte(response),
			})
			if err != nil {
				logger.Error("Failed to publish a message.", err)
			}
			if err := d.Ack(false); err != nil { // Acknowledge a message maunally.
				logger.Error("Failed to acknowledge a message.", err)
			}
		}
	}()
	fmt.Println(" [*] Awaiting RPC requests")
	<-forever
}