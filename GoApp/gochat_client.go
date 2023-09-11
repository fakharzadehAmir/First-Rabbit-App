package main

import (
	// "bufio"
	"fmt"
	// "os"

	"github.com/streadway/amqp"
)

func main() {
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

	// _, err = ch.QueueDeclare(
	// 	"go_sender",
	// 	false,
	// 	false,
	// 	false,
	// 	false,
	// 	nil,
	// )
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Print("What's your name? ")
	// scanner.Scan()
	// your_name := scanner.Text()

	// my_turn := true

	fmt.Println("Wait")

	// for {
	// 	if my_turn {
	// 		fmt.Print("You: ")
	// 		scanner.Scan()
	// 		your_message := scanner.Text()
	// 		err = ch.Publish(
	// 			"", "go_sender", false, false,
	// 			amqp.Publishing{
	// 				ContentType: "text/plain",
	// 				Body:        []byte(your_name + ": " + your_message),
	// 			})
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}

	// 	} else {
	// 		messages, _ := ch.Consume(
	// 			"py_sender", "", true, false, false, false, nil,
	// 		)
	// 		for message := range messages {
	// 			fmt.Printf("%s\n", message.Body)
	// 		}
	// 	}
	// 	my_turn = !my_turn

	// }

	messages, _ := ch.Consume(
		"py_sender", "", true, false, false, false, nil,
	)
	recieve := make(chan bool)
	go func() {
		for message := range messages {
			fmt.Printf("%s\n", message.Body)
		}
	}()
	<-recieve

}
