package main

import (
	"fmt"
	"github.com/xander/economic_cdn/classes"
)

func main() {
	fmt.Println("Hello, World!")

	// Create a new MessageReceiver instance
	messageReceiver := classes.NewMessageReceiver(8080)

	// Start listening for incoming messages concurrently
	go func() {
		if err := messageReceiver.StartListening(); err != nil {
			fmt.Println(err)
		}
	}()

	// Create a new MessageSender instance
	messageSender := classes.NewMessageSender("http://localhost:8080")

	// Message to be sent
	message := "Hello, this is a test message!"

	// Send the message using the MessageSender
	if err := messageSender.SendMessage(message); err != nil {
		fmt.Println(err)
	}

	// Sleep for a while to allow the receiver to process the message
	// This is a simplistic approach, and you may need more sophisticated synchronization
	// mechanisms in a real-world scenario.
	// For example, you might use channels to signal the completion of message processing.
	select {}
}
