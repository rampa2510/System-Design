package commons

import "fmt"

func AcceptInput(message string, inputVariable *int) {
	fmt.Print(message)
	_, err := fmt.Scanf("%d", inputVariable)

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
}
