package loadbalancingalgo

import (
	"fmt"

	"github.com/rampa2510/system-design/commons"
)

func ConsistentHashing() {
	var serverRing []int
	var numberOfInitialServers int

	commons.AcceptInput("Enter number of initial servers (0 to exit):", &numberOfInitialServers)

	if numberOfInitialServers == 0 {
		fmt.Println("Exiting consistent hashing alg. GoodBye!")
		return
	}

	for i := 0; i <= numberOfInitialServers; i++ {
		serverRing = append(serverRing, 1)
	}

	for i, v := range serverRing {
		message := fmt.Sprintf("index %d value %d", i, v)
		fmt.Println(message)
	}

}
