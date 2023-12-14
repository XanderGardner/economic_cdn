package main

import (
	"bytes"
	"fmt"
	"net/http"
)

// user class for sending messages to a specific port
type MessageSender struct {
	URL string
}

// NewMessageSender creates a new instance of MessageSender with the given URL.
func NewMessageSender(url string) *MessageSender {
	return &MessageSender{URL: url}
}

// SendMessage sends a message to the specified URL.
func (ms *MessageSender) SendMessage(message string) error {
	resp, err := http.Post(ms.URL, "text/plain", bytes.NewBufferString(message))
	if err != nil {
		return fmt.Errorf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	fmt.Println("Message sent successfully!")
	return nil
}

func main() {
	fmt.Println("Hello from folder1/file1!")
}


