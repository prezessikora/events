package fanout

import (
	"context"
	"github.com/prezessikora/events/errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func Connect() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	if err != nil {
		errors.LogError(err, "could not connect to rabbit mq")
		return nil, err
	}
	return conn, nil
}

func CloseConnection(conn *amqp.Connection) {
	err := conn.Close()
	if err != nil {
		log.Printf("error closing the rabbit mq connection: w%v", err)
	}

}

// Subscribe creates exclusive queue bound to the fanout exchange and returns channel for incoming messages
func Subscribe(exchangeName string, process func([]byte)) error {
	conn, err := Connect()
	if err != nil {
		return err
	}

	defer CloseConnection(conn)

	ch, err := conn.Channel()
	if err != nil {
		errors.LogError(err, "failed to open rabbit channel")
		return err
	}
	defer ch.Close()

	err = DeclareFanoutExchange(ch, exchangeName)
	if err != nil {
		errors.LogError(err, "failed to declare exchange")
		return err
	}

	q, err := ch.QueueDeclare( // anonymous queue
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		errors.LogError(err, "failed to declare queue")
		return err
	}

	err = ch.QueueBind(
		q.Name,       // queue name
		"",           // routing key
		exchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		errors.LogError(err, "failed to bind queue to exchange")
		return err
	}
	msgChannel, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		errors.LogError(err, "failed to register queue to exchange")
		return err
	}

	var forever chan struct{}

	go func() {
		for d := range msgChannel {
			process(d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

	return nil
}

func PublishFanOut(exchangeName string, message string) error {
	conn, err := Connect()
	if err != nil {
		return err
	}
	defer CloseConnection(conn)

	ch, err := conn.Channel()
	if err != nil {
		errors.LogError(err, "failed to open rabbit channel")
		return err
	}
	defer ch.Close()

	err = DeclareFanoutExchange(ch, exchangeName)

	if err != nil {
		errors.LogError(err, "failed to create fanout exchange")
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Printf("Publishing event with body: [%v] \n", message)
	err = ch.PublishWithContext(ctx,
		exchangeName, // exchange
		"",           // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

	if err != nil {
		errors.LogError(err, "failed to create fanout exchange")
		return err
	}

	return nil
}

func DeclareFanoutExchange(ch *amqp.Channel, exchangeName string) error {
	err := ch.ExchangeDeclare(
		exchangeName, // name
		"fanout",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	return err
}
