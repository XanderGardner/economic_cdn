package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

// OriginServer is a representation of an origin server listening at a port
// and serving requests
type OriginServerStats struct {
	Port int
}

// NewOriginServerStats creates a new instance of OriginServerStats with the given port
func NewOriginServerStats(port int) *OriginServerStats {
	return &OriginServerStats{
		Port: port,
	}
}

// StartListening starts listens for incoming messages.
func (mr *OriginServerStats) StartListening() error {
	// Simple HTTP handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		
		// handle the incoming message
		fmt.Printf("Received message: %s\n", body)


	}

	// Start the server on the specified port
	http.HandleFunc("/", handler)
	fmt.Printf("Server listening on :%d\n", mr.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", mr.Port), nil)
}

func main() {
	// Default port
	default_port := 8080

	// Check if a port is provided as a command-line argument
	if len(os.Args) > 1 {
		port1, err1 := strconv.Atoi(os.Args[1])
		if err1 == nil{
			default_port = port1
		}
	}

	// Create an instance of OriginServer with the specified port
	receiver := NewOriginServerStats(default_port)

	// Start listening for incoming messages
	err := receiver.StartListening()
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
