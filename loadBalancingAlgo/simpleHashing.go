package loadbalancingalgo

import (
	"fmt"
	"github.com/rampa2510/system-design/commons"
)

func getServerNumber(noOfServers, clientId int) int {
	return (clientId % noOfServers) + 1
}

func SimpleHashing() {
	var numberOfServers int
	var clientId int

	commons.AcceptInput("Enter number of servers: ", &numberOfServers)

	for {
		commons.AcceptInput("Enter clientId (0 to exit):", &clientId)

		if clientId == 0 {
			fmt.Println("Exiting simple hasing algo. Goodbye!")
			return
		}

		message := fmt.Sprintf("The request with %d will be route to server %d", clientId, getServerNumber(numberOfServers, clientId))
		fmt.Println(message)
	}
}
