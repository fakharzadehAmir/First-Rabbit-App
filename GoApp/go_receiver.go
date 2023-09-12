package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func goReceiver() {
	connection, err := amqp.Dial("amqp://chat_app:chat_app_password@localhost:5672/")
	if err != nil {
		fmt.Println("Failed Connecting to RabbitMQ")
		panic(err)
	}
	ch, err := connection.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ch.Close()
	messages, _ := ch.Consume(
		"py_sender", "", true, false, false, false, nil,
	)
	for message := range messages {
		fmt.Printf("%s\n", message.Body)
	}
	connection.Close()

}
