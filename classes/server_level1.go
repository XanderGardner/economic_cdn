package classes

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// MessageReceiver is a simple class for receiving messages on a specific port.
type MessageReceiver struct {
	Port int
}

// NewMessageReceiver creates a new instance of MessageReceiver with the given port.
func NewMessageReceiver(port int) *MessageReceiver {
	return &MessageReceiver{Port: port}
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

		fmt.Printf("Received message: %s\n", body)
	}

	// Start the server on the specified port
	http.HandleFunc("/", handler)
	fmt.Printf("Server listening on :%d\n", mr.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", mr.Port), nil)
}
