package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"github.com/xander/economic_cdn/caches"
	// "github.com/xander/economic_cdn/conversion"
)


// MessageReceiver is a simple class for receiving messages on a specific port.
type MessageReceiver struct {
	ServerName string
	UserPort int
	Level2Port int
	StatsPort int
	ServerCache cache.Cache 
	CurrentRequestCount int
}

// NewMessageReceiver creates a new instance of MessageReceiver with the given port.
func NewMessageReceiver(server_name string, user_port int, level2_port int, stats_port int) *MessageReceiver {
	LEVEL1_CACHE_SIZE := 5000

	return &MessageReceiver{
		ServerName: server_name,
		UserPort: user_port,
		Level2Port: level2_port,
		StatsPort: stats_port,
		ServerCache: cache.NewFifo(LEVEL1_CACHE_SIZE),
		CurrentRequestCount: 0,
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
		key_requested := string(body)

		// handle the incoming message
		fmt.Printf("Received message: %s\n", body)

		// check if the body is in the cache
		byte_val, ok := mr.ServerCache.Get(key_requested)

		if ok {
			// already in current cache
			fmt.Printf("    Cache Hit\n")
			
			// respond to user
			w.Write(byte_val)
		} else {
			// not in current cache
			fmt.Printf("    Cache Miss\n")

			// request from level 2
			level2_response, _ := mr.SendMessage(key_requested)

			// respond to user
			w.Write([]byte(level2_response))

			// add to cache
			mr.ServerCache.Set(key_requested, []byte(level2_response))

		}

		// handle the increased number of requests if needed
		mr.CurrentRequestCount += 1
		if mr.CurrentRequestCount > 10 {
			mr.SendStatUpdate()
			mr.CurrentRequestCount = 0
		}

		
	}

	// Start the server on the specified port
	http.HandleFunc("/", handler)
	fmt.Printf("Server listening on :%d\n", mr.UserPort)
	return http.ListenAndServe(fmt.Sprintf(":%d", mr.UserPort), nil)
}

// SendMessage sends a message to the specified port.
func (mr *MessageReceiver) SendMessage(message string) (string, error) {
	url := fmt.Sprintf("http://localhost:%d", mr.Level2Port)
	resp, err := http.Post(url, "text/plain", strings.NewReader(message))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Error sending message. Status: %s", resp.Status)
	}

	// Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading response body: %v", err)
	}

	return string(responseBody), nil
}

// SendMessage sends a message to the specified port.
func (mr *MessageReceiver) SendStatUpdate() {
	url := fmt.Sprintf("http://localhost:%d", mr.StatsPort)
	http.Post(url, "text/plain", strings.NewReader(mr.ServerName))
}











func main() {
	// Default port
	default_user_port := 8080
	default_level2_port := 8081
	default_stats_port := 8087
	default_name := "unknown_server"

	// Check if a port is provided as a command-line argument
	if len(os.Args) > 4 {
		port1, err1 := strconv.Atoi(os.Args[1])
		port2, err2 := strconv.Atoi(os.Args[2])
		port3, err3 := strconv.Atoi(os.Args[3])
		if err1 == nil && err2 == nil && err3 == nil {
			default_user_port = port1
			default_level2_port = port2
			default_stats_port = port3
			default_name = os.Args[4]
		}
	} else {
		return
	}

	// Create an instance of MessageReceiver with the specified port
	receiver := NewMessageReceiver(default_name, default_user_port, default_level2_port, default_stats_port)

	// Start listening for incoming messages
	err := receiver.StartListening()
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
