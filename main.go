package main

import (
	"fmt"
	"github.com/rampa2510/system-design/commons"
	"github.com/rampa2510/system-design/loadBalancingAlgo"
)

type AlgorithmFunc func()

var AlgorithmMap = map[int]struct {
	Name string
	Func AlgorithmFunc
}{
	1: {"Simple Hashing", loadbalancingalgo.SimpleHashing},
}

func main() {
	for {
		fmt.Println("\nSelect an algorithm to run:")
		for num, algo := range AlgorithmMap {
			fmt.Printf("%d. %s\n", num, algo.Name)
		}

		var choice int
		commons.AcceptInput("Enter your choice (0 to exit): ", &choice)

		if choice == 0 {
			fmt.Println("Exiting the program. Goodbye!")
			return
		}

		if algo, ok := AlgorithmMap[choice]; ok {
			fmt.Printf("\nRunning %s\n", algo.Name)
			algo.Func()
		} else {
			fmt.Println("Invalid selection. Please try again.")
		}
	}

}
