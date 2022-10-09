package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/joho/godotenv/autoload"
	amqp "github.com/rabbitmq/amqp091-go"
)

const ROUTING_KEY_POST = ".message"

func main() {

	amqpConnection := createConnection()

	defer amqpConnection.Close()
	_, err := createQueue(amqpConnection, "taxi.1", "taxi-direct")

	if err != nil {
		fmt.Println("Cannot create queue", err)
	} else {

	}
	// taxiQueue.Messages

	router := chi.NewRouter()
	router.Use()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8090"
	}
	log.Println("** About starting server on Port " + port + " **")

	//All http endpoints should be loaded in the init function of router.go
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal("Error with http server", err)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func orderTaxi(conn *amqp.Connection, routingKey, exchangeName string) (err error) {
	channel, err := conn.Channel()

	if err != nil {
		log.Println("cannot create channel for order taxi", err)
		return
	}
	channel.Confirm(true)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	channel.PublishWithContext(ctx, // The context
		exchangeName, // The exchange to send message to
		routingKey,   // The routing key to match the queue binding
		false,        // mandatory: to ensure the exchange and queue exists, else errors
		false,        // immediate: ensures immediate delivery to a consumer i.e places it at the top of the queue
		amqp.Publishing{
			Body:            []byte("I am booking taxi"),
			ContentEncoding: "UTF-8",
			ContentType:     "application/json",
			MessageId:       "45", // random number

		},
	)

	defer channel.Close()

	return
}

func createQueue(conn *amqp.Connection, name, exchangeName string) (qu *amqp.Queue, err error) {

	// create channel
	channel, err := conn.Channel()

	if err != nil {
		log.Println("cannot create channel. Error: " + err.Error())
		return
	}

	// create Exchange
	err = channel.ExchangeDeclare(exchangeName, //name
		"direct", // type of exchange The common types are "direct", "fanout", "topic" and "headers".
		true,     // durable
		false,    // auto-delete
		false,    // internal exchange
		false,    // no-wait
		nil,      // arguments
	)

	if err != nil {
		log.Println("cannot create exchange" + exchangeName + ". Error: " + err.Error())
		return
	}

	// create queue

	/* create a queue for taxi driver 1

		NOTE: durable means to survive when server restarts
			   auto-deleted means queue will be deleted when last consumer is
			   canceled or the last consumer's channel is closed.

		   Durable and Non-Auto-Deleted queues (grouped as one) will survive server restarts and remain
		   when there are no remaining consumers or bindings.

		   Non-Durable and Auto-Deleted queues (grouped as one) will not be redeclared on server restart
		   and will be deleted by the server after a short time when the last consumer is
	       canceled or the last consumer's channel is closed.

		   Exclusive queues are only accessible by the connection that declares them and
		   will be deleted when the connection closes.
	*/
	// Declare a queue for a given taxi
	queue1, err := channel.QueueDeclare(name, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		log.Println("cannot create queue " + name + ". Error: " + err.Error())
		return
	}

	// Bind the queue to the exchange
	err = channel.QueueBind(
		queue1.Name,                  //The queue name
		queue1.Name+ROUTING_KEY_POST, // The binding key. I am using the format QueueName.message
		exchangeName,
		false,
		nil,
	)

	if err != nil {
		log.Println("error cannot bind ", queue1.Name, " to exchange ", exchangeName)
		return
	}

	return &queue1, nil
}
func createConnection() *amqp.Connection {
	rabbitUrl := os.Getenv("RABBITMQ_URI")
	// fmt.Println(rabbitUrl, "Rabbit URL")
	conn, err := amqp.Dial(rabbitUrl)
	// amqp.DialConfig()

	failOnError(err, "Failed to connect to RabbitMQ")

	return conn
}
