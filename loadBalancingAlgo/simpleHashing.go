package main

import (
	"fmt"
	"github.com/rampa2510/system-design/commons"
)

func getServerNumber(noOfServers, clientId int) int {
	return clientId % noOfServers
}

func main() {
	var numberOfServers int
	var clientId int

	commons.AcceptInput("Enter number of servers: ", &numberOfServers)

	for {
		commons.AcceptInput("Enter clientId :", &clientId)
		message := fmt.Sprintf("The request with %d will be route to server %d", clientId, getServerNumber(numberOfServers, clientId))
		fmt.Println(message)
	}
}
