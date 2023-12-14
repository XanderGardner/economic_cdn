package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
	"container/list"
	"sort"
)

// OriginServer is a representation of an origin server listening at a port
// and serving requests
type OriginServerStats struct {
	Port int
	StartTime time.Time
	Signals map[string]*list.List
}

// NewOriginServerStats creates a new instance of OriginServerStats with the given port
func NewOriginServerStats(port int) *OriginServerStats {
	return &OriginServerStats{
		Port: port,
		StartTime: time.Now(),
		Signals: make(map[string]*list.List),
	}
}

// StartListening starts listens for incoming messages.
func (mr *OriginServerStats) StartListening() error {
	// Simple HTTP handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		currTime := mr.getSecondTime()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		
		// handle the incoming message
		fmt.Printf("Received message: %s\n", body)
		server_signaled := string(body)

		// when a server signals that they received another 10 requests, update that servers signal queue
		curr_list, ok := mr.Signals[server_signaled]
		
		if ok {
			curr_list.PushBack(currTime)
			if curr_list.Len() > 5 {
				curr_list.Remove(curr_list.Front())
			}
		} else {
			new_list := list.New()
			new_list.PushBack(currTime)
			mr.Signals[server_signaled] = new_list
		}

		// body conains just a string with the name of the server that sent the request meaning that that server got 10 messages. 
		// now we can update the time for all the servers

		mr.PrintUpdatedStats(currTime)


	}

	// Start the server on the specified port
	http.HandleFunc("/", handler)
	fmt.Printf("Server listening on :%d\n", mr.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", mr.Port), nil)
}

// gets number of sceonds since the OriginServerStats was created
func (mr *OriginServerStats) getSecondTime() float64 {
	elapsed := time.Since(mr.StartTime)
	return elapsed.Seconds()
}

// for each server string in our map, print average requests per second for the past 5 signals
// func (mr *OriginServerStats) PrintUpdatedStats(currTime float64) {
// 	fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
// 	fmt.Printf("Server   |   Current AVG Requests Per Second \n\n")

// 	for server, signal_list := range mr.Signals {
// 		oldestTime := signal_list.Front().Value.(float64)
// 		avg_requests_per_second := 50 / (currTime - oldestTime)
// 		fmt.Printf("%s      |   %v\n", server, avg_requests_per_second)
// 	}

// }

type ServerData struct {
	Server              string
	AvgRequestsPerSecond float64
}

func (mr *OriginServerStats) PrintUpdatedStats(currTime float64) {
	fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Printf("Server   |   Current AVG Requests Per Second \n\n")

	// Create a slice to hold the server data
	var serverData []ServerData

	// Iterate over servers and calculate average requests per second
	for server, signalList := range mr.Signals {
		oldestTime := signalList.Front().Value.(float64)
		avgRequestsPerSecond := 50 / (currTime - oldestTime)

		// Append server data to the slice
		serverData = append(serverData, ServerData{Server: server, AvgRequestsPerSecond: avgRequestsPerSecond})
	}

	// Sort server data by AvgRequestsPerSecond in descending order
	sort.Slice(serverData, func(i, j int) bool {
		return serverData[i].AvgRequestsPerSecond > serverData[j].AvgRequestsPerSecond
	})

	// Print sorted data
	for _, data := range serverData {
		fmt.Printf("%s | %.2f requests per second\n", data.Server, data.AvgRequestsPerSecond)
	}
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
