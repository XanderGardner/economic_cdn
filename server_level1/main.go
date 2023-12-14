package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// MessageReceiver is a simple class for receiving messages on a specific port.
type MessageReceiver struct {
	UserPort int
	Level2Port int
}

// NewMessageReceiver creates a new instance of MessageReceiver with the given port.
func NewMessageReceiver(user_port int, level2_port int) *MessageReceiver {
	return &MessageReceiver{
		UserPort: user_port,
		Level2Port: level2_port,
	}
}

// StartListening starts the message receiver and listens for incoming messages.
func (mr *MessageReceiver) StartListening() error {
	// Simple HTTP handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		// handle the incoming message
		fmt.Printf("Received message: %s\n", body)
		mr.SendMessage(string(body))
	}

	// Start the server on the specified port
	http.HandleFunc("/", handler)
	fmt.Printf("Server listening on :%d\n", mr.UserPort)
	return http.ListenAndServe(fmt.Sprintf(":%d", mr.UserPort), nil)
}

// SendMessage sends a message to the specified port.
func (mr *MessageReceiver) SendMessage(message string) error {
	url := fmt.Sprintf("http://localhost:%d", mr.Level2Port)
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
	default_user_port := 8080
	default_level2_port := 8081

	// Check if a port is provided as a command-line argument
	if len(os.Args) > 2 {
		port1, err1 := strconv.Atoi(os.Args[1])
		port2, err2 := strconv.Atoi(os.Args[2])
		if err1 == nil && err2 == nil {
			default_user_port = port1
			default_level2_port = port2
		}
	}

	// Create an instance of MessageReceiver with the specified port
	receiver := NewMessageReceiver(default_user_port, default_level2_port)

	// Start listening for incoming messages
	err := receiver.StartListening()
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
