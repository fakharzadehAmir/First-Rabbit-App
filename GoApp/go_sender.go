package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

func goSender() {
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
	fmt.Print("You: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	your_message := scanner.Text()
	err = ch.Publish(
		"", "go_sender", false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("He said: " + your_message),
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	connection.Close()

}
