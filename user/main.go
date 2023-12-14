// main.go
package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"io/ioutil"
	"unicode"
	"time"
	"math/rand"
)





//////// MessageSender (user)

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

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error reading response body: %s\n", err)
	}

	// Print the response content
	fmt.Printf("    Got CDN response: " + string(body) + "\n")

	return nil
}




//////// helper functions

// Given a string, returns the strings with only characters in the alphabete and spaces
func keepAlphabeticAndSpaces(input string) string {
	var result []rune

	for _, char := range input {
		// Keep alphabetic characters and spaces
		if unicode.IsLetter(char) || unicode.IsSpace(char) {
			result = append(result, char)
		}
	}

	return string(result)
}

// Pauses the process execution for a random duration around one second
func randomPause() {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate a random duration between 0.5 and 1.5 seconds
	minDuration := 500 * time.Millisecond
	maxDuration := 1500 * time.Millisecond
	randomDuration := time.Duration(rand.Int63n(int64(maxDuration-minDuration)) + int64(minDuration))

	// Pause for the generated random duration
	time.Sleep(randomDuration)
}




//////// main function

func main() {
	// Default port
	defaultPort := 8080
	message_text := "hello!"

	// Check if a port is provided as a command-line argument
	if len(os.Args) > 1 {
		port, err := strconv.Atoi(os.Args[1])
		if err == nil {
			defaultPort = port
		}
	}
	
	if len(os.Args) > 2 {
		// Read the contents of the file
		filePath := os.Args[2]
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file: %s\n", err)
			return
		}

		// Convert the byte slice to a string
		message_text = keepAlphabeticAndSpaces(string(content))

	}

	// Create an instance of MessageSender with the specified port
	sender := NewMessageSender(defaultPort)

	

	// Iterate over the words and send request to the cdn
	words := strings.Split(message_text, " ")
	for _, word := range words {
		
		// Send the word to the cdn (request the value, which is the lenght of the string)
		fmt.Printf("Reqesting \"" + word + "\" from CDN" + "\n")
		err := sender.SendMessage(word)
		if err != nil {
			fmt.Printf("Error sending message: %s\n", err)
		}
		
		// Pause for about a second
		randomPause()


	}
	// Send a message
	


	
}
