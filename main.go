package main

import (
	"github.com/streadway/amqp"
)

func main() {
	url := "amqp://guest:guest@localhost:5672"

	conn, err := amqp.Dial(url)

	defer conn.Close()

	ch, err := conn.Channel()

	err = ch.ExchangeDeclare("events", "topic", true, false, false, false, nil)
	failOnError(err)

	message := amqp.Publishing{
		Body: []byte("Hello world"),
	}

	err = ch.Publish("events", "random-key", false, false, message)

	_, err = ch.QueueDeclare("test", true, false, false, false, nil)

	msgs, err := ch.Consume("test", "", false, false, false, false, nil)

	failOnError(err)

	for msg := range msgs {
		println(string(msg.Body))
		msg.Ack(false)
	}
}

func failOnError(e error) {
	if e != nil {
		panic(e.Error())
	}
}
