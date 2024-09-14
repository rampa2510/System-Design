package loadbalancingalgo

import (
	"fmt"
	"github.com/rampa2510/system-design/commons"
	"math"
	"math/rand"
)

type Server struct {
	id          int
	connections int
}

func simulateInitialLoad(servers []Server) {
	for i := range servers {
		servers[i].connections = rand.Intn(5) // Random initial load between 0 and 4
	}
}

func findLeastLoadedServer(servers []Server) (int, error) {
	if len(servers) == 0 {
		return -1, fmt.Errorf("no servers available")
	}

	minConnections := math.MaxInt32
	leastLoadedServer := -1

	for _, server := range servers {
		if server.connections < minConnections {
			minConnections = server.connections
			leastLoadedServer = server.id
		}
	}

	if leastLoadedServer == -1 {
		return -1, fmt.Errorf("all servers are at maximum capacity")
	}

	return leastLoadedServer, nil
}

func printSystemState(serverNodes []Server) {
	for _, server := range serverNodes {
		fmt.Printf("Server %d: %d connections\n", server.id, server.connections)
	}

}

func LeastConnection() {
	var noOfServers int
	commons.AcceptInput("Enter no of servers (0 to exit): ", &noOfServers)
	if noOfServers == 0 {
		fmt.Println("Exiting Least Connection. Goodbye!")
		return
	}

	var noOfRequests int
	commons.AcceptInput("Enter no of requests (0 to exit): ", &noOfRequests)
	if noOfRequests == 0 {
		fmt.Println("Exiting Least Connection. No requests to process.")
		return
	}

	servers := make([]Server, noOfServers)
	for i := range servers {
		servers[i] = Server{id: i}
	}

	simulateInitialLoad(servers)

	fmt.Println("Initial server states:")
	printSystemState(servers)

	fmt.Println("\nProcessing requests...")
	for i := 0; i < noOfRequests; i++ {
		selectedServer, err := findLeastLoadedServer(servers)
		if err != nil {
			fmt.Printf("Error processing request %d: %v\n", i+1, err)
			continue
		}

		servers[selectedServer].connections++
		fmt.Printf("Request %d assigned to server %d (now has %d connections)\n",
			i+1, selectedServer, servers[selectedServer].connections)
	}

	fmt.Println("\nFinal server states:")

	printSystemState(servers)
}
