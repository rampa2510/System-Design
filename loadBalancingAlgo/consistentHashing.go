package loadbalancingalgo

import (
	"fmt"
	"github.com/rampa2510/system-design/commons"
)

func addServer(serverRing *[]int) {
	*serverRing = append(*serverRing, 1)
}

func removeServer(serverRing *[]int) {
	lenOfServerRing := len(*serverRing)
	if lenOfServerRing > 0 {
		*serverRing = (*serverRing)[:lenOfServerRing-1]
	}
}

func getServerForClientId(clientId int, serverRing []int) int {
	lenOfServerRing := len(serverRing)
	if lenOfServerRing == 0 {
		return -1
	}
	initialServerIdx := clientId % lenOfServerRing
	for i := 0; i < lenOfServerRing; i++ {
		idx := (initialServerIdx + i) % lenOfServerRing
		if serverRing[idx] == 1 {
			return idx
		}
	}
	return -1
}

func printServerRing(serverRing []int) {
	for i, v := range serverRing {
		fmt.Printf("At index %d, server state is %d\n", i, v)
	}
}

func ConsistentHashing() {
	var serverRing []int
	var numberOfInitialServers int

	commons.AcceptInput("Enter number of initial servers (0 to exit):", &numberOfInitialServers)
	if numberOfInitialServers == 0 {
		fmt.Println("Exiting consistent hashing alg. GoodBye!")
		return
	}

	for i := 0; i < numberOfInitialServers; i++ {
		serverRing = append(serverRing, 1)
	}

	for {
		fmt.Println("\nChoose an action:")
		fmt.Println("1. Add Server at the end")
		fmt.Println("2. Remove Server at the end")
		fmt.Println("3. Get Server for Client")
		fmt.Println("0. Exit")

		var actionChoice int
		commons.AcceptInput("Enter action choice:", &actionChoice)

		switch actionChoice {
		case 0:
			fmt.Println("Exiting Consistent hashing program. Goodbye!")
			return
		case 1:
			addServer(&serverRing)
			printServerRing(serverRing)
		case 2:
			removeServer(&serverRing)
			printServerRing(serverRing)
		case 3:
			var clientId int
			commons.AcceptInput("Enter Client Id: ", &clientId)
			serverIdx := getServerForClientId(clientId, serverRing)
			fmt.Printf("Client with id %d will be assigned server with idx %d.\n", clientId, serverIdx)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

