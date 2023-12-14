// main.go
package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// MessageSender is a simple class for sending messages to a specific port.
type MessageSender struct {
	Port int
}

// NewMessageSender creates a new instance of MessageSender with the given port.
func NewMessageSender(port int) *MessageSender {
	return &MessageSender{Port: port}
}

// SendMessage sends a message to the specified port.
func (ms *MessageSender) SendMessage(message string) error {
	url := fmt.Sprintf("http://localhost:%d", ms.Port)
	resp, err := http.Post(url, "text/plain", strings.NewReader(message))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Error sending message. Status: %s", resp.Status)
	}

	return nil
}

func main() {
	// Default port
	defaultPort := 8080

	// Check if a port is provided as a command-line argument
	if len(os.Args) > 1 {
		port, err := strconv.Atoi(os.Args[1])
		if err == nil {
			defaultPort = port
		}
	}

	// Create an instance of MessageSender with the specified port
	sender := NewMessageSender(defaultPort)

	// Send a message
	message := "Hello from user!"
	err := sender.SendMessage(message)
	if err != nil {
		fmt.Printf("Error sending message: %s\n", err)
	}
}
