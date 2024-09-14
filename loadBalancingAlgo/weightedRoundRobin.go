package loadbalancingalgo

import (
	"fmt"

	"github.com/rampa2510/system-design/commons"
)

type ServerNode struct {
	weight      int
	connections int
}

func getSystemState(serverNodes []ServerNode) {

	noOfServers := len(serverNodes)
	for i := 0; i < noOfServers; i++ {
		server := serverNodes[i]
		fmt.Printf("Server at index %d has weight %d and live connections %d\n", i, server.weight, server.connections)
	}
}

func serveRequests(serverNodes *[]ServerNode) {

	noOfServers := len(*serverNodes)

	for i := 0; i < noOfServers; i++ {

		if (*serverNodes)[i].weight != (*serverNodes)[i].connections {
			(*serverNodes)[i].connections++
			return
		}
	}

}

func WeightedRoundRobin() {

	var noOfServers int
	commons.AcceptInput("Enter Number of servers (0 to exit): ", &noOfServers)

	if noOfServers == 0 {
		fmt.Println("Exiting Weighted Round Robin. Goodbye!")
		return
	}

	serverNodes := make([]ServerNode, noOfServers)

	var weight int
	for i := 0; i < noOfServers; i++ {
		inputMsg := fmt.Sprintf("Enter weight for server at idx %d: ", i+1)

		commons.AcceptInput(inputMsg, &weight)
		serverNodes[i].weight = weight
	}

	var noOfRequests int
	commons.AcceptInput("Enter no of requests (o to exit): ", &noOfRequests)

	getSystemState(serverNodes)
	fmt.Println("Running algo..........")
	for i := 0; i < noOfRequests; i++ {
		serveRequests(&serverNodes)
	}
	getSystemState(serverNodes)

}
