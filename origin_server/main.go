package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// OriginServer is a representation of an origin server listening at a port
// and serving requests
type OriginServer struct {
	Level2Port int
	Database map[string]int
}

// NewOriginServer creates a new instance of OriginServer with the given port
func NewOriginServer(level2_port int) *OriginServer {
	return &OriginServer{
		Level2Port: level2_port,
		Database: make(map[string]int),
	}
}

// StartListening starts listens for incoming messages.
func (mr *OriginServer) StartListening() error {
	// Simple HTTP handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		// handle the incoming message
		fmt.Printf("Received message: %s\n", body)

		// respond to user
		value, exists := mr.Database[string(body)]
		if exists {
			response := strconv.Itoa(value)
			w.Write([]byte(response))
		} else {
			mr.Database[string(body)] = len(string(body))
			response := strconv.Itoa(len(string(body)))
			w.Write([]byte(response))
		}

	}

	// Start the server on the specified port
	http.HandleFunc("/", handler)
	fmt.Printf("Server listening on :%d\n", mr.Level2Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", mr.Level2Port), nil)
}

// SendMessage sends a message to the specified port.
func (mr *OriginServer) SendMessage(message string) error {
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
	default_level2_port := 8080

	// Check if a port is provided as a command-line argument
	if len(os.Args) > 1 {
		port1, err1 := strconv.Atoi(os.Args[1])
		if err1 == nil{
			default_level2_port = port1
		}
	}

	// Create an instance of OriginServer with the specified port
	receiver := NewOriginServer(default_level2_port)

	// Start listening for incoming messages
	err := receiver.StartListening()
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
