package main

import (
	"fmt"
	"main/messagereceiver"
)

func main() {
	fmt.Println("Hello, World!")

	// Create a new MessageReceiver instance
	messageReceiver := NewMessageReceiver(8080)

	// Start listening for incoming messages
	if err := messageReceiver.StartListening(); err != nil {
		fmt.Println(err)
	}

	// Create a new MessageSender instance
	messageSender := NewMessageSender("http://localhost:8080")

	

	// Message to be sent
	message := "Hello, this is a test message!"

	// Send the message using the MessageSender
	if err := messageSender.SendMessage(message); err != nil {
		fmt.Println(err)
	}
}
