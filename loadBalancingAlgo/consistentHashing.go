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

	(*serverRing)[lenOfServerRing-1] = 0

}

func printServerRing(serverRing []int) {
	for i, v := range serverRing {
		fmt.Printf("At index %d, server state is %d\n", i, v)
	}
}

func ConsistentHashing() {
	var serverRing []int
	var numberOfInitialServers int

	type actionFunc func(*[]int)

	var actionMap = map[int]struct {
		actionType string
		function   actionFunc
	}{
		1: {"Add Server at the end", addServer},
		2: {"Remove Server at the end", removeServer},
	}

	commons.AcceptInput("Enter number of initial servers (0 to exit):", &numberOfInitialServers)

	if numberOfInitialServers == 0 {
		fmt.Println("Exiting consistent hashing alg. GoodBye!")
		return
	}

	for i := 0; i <= numberOfInitialServers; i++ {
		serverRing = append(serverRing, 1)
	}

	// for i, v := range serverRing {
	// 	message := fmt.Sprintf("index %d value %d", i, v)
	// 	fmt.Println(message)
	// }

	for {
		for num, action := range actionMap {
			fmt.Printf("%d. %s\n", num, action.actionType)
		}

		var actionChoice int
		commons.AcceptInput("Enter action choice (0 to exit):", &actionChoice)

		if actionChoice == 0 {
			fmt.Println("Exiting Consistent hashing program. Goodbye!")
			return
		}

		if action, ok := actionMap[actionChoice]; ok {
			action.function(&serverRing)
		}

		if actionChoice < 3 {
			printServerRing(serverRing)

		}
	}

}
