package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
	"github.com/LiamYabou/top100-ranking/app"
	"github.com/LiamYabou/top100-ranking/preference"
	"github.com/LiamYabou/top100-ranking/variable"
	"github.com/LiamYabou/top100-pkg/logger"
	"github.com/LiamYabou/top100-ranking/api"
	"github.com/streadway/amqp"
	"github.com/panjf2000/ants/v2"
)

func main() {
	defer app.Finalize()
	c, _ := strconv.Atoi(variable.Concurrency)
	opts := &preference.Options{
		DB: app.DBpool,
		AMQP: app.AMQPconn,
		Concurrency:   c,
		PrefetchCount: (c * 2),
		InvokerInterval: 200,
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
		opts.PrefetchCount,     // prefetch count
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
	fmt.Println(" [*] Awaiting RPC requests")
	var wg sync.WaitGroup
	opts = preference.LoadOptions(preference.WithOptions(*opts), preference.WithDelivery(msgs))
	p, _ := ants.NewPoolWithFunc(opts.Concurrency, func(optionsInterface interface{}) {
		opts, ok := optionsInterface.(*preference.Options)
		if !ok {
			logger.Error("The type `*preference.Options` has not implemented the interface `optionsInterface`.", nil)
		}
		for d := range opts.Delivery {
			args := strings.Split(string(d.Body), "/")
			// args[0] represents the action of the consumer;
			// args[1] represents the id of the category row;
			// args[2] represents the number of the page.
			// `performDispatcher` method dispaches the workers to perform the specific tasks according to the argument of the action.
			performDispatcher(ch, d, opts, args)
		}
		wg.Done()
	})
	defer p.Release()
	for {
		wg.Add(1)
		_ = p.Invoke(opts)
		// set the interval of invoking work.
		time.Sleep(time.Duration(opts.InvokerInterval) * time.Millisecond)
	}
}

func performDispatcher(ch *amqp.Channel, delivery amqp.Delivery, opts *preference.Options, args []string) {
	categoryId, _ := strconv.Atoi(args[1]) 
	page, _ := strconv.Atoi(args[2])
	response := api.FindRankings(categoryId, page, opts)
	err := ch.Publish(
		"",        // exchange
		delivery.ReplyTo, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: delivery.CorrelationId,
			Body:          []byte(response),
	})
	if err != nil {
		logger.Error("Failed to publish a message.", err)
	}
	if err := delivery.Ack(false); err != nil { // Acknowledge a message maunally.
		logger.Error("Failed to acknowledge a message.", err)
	}
}
