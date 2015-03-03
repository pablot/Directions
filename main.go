package main

import "fmt"

func menu() {
	fmt.Println("Choose:")
	fmt.Println("1. Easy")
	fmt.Println("2. Medium")
	fmt.Println("3. Hard")
	fmt.Println("Other value. Quit")
}

func getResponse() int {
	var response int
	fmt.Scanf("%d", &response)
	return response
}

func game(mode int) {
}

func main() {
	menu()
	var response = getResponse()
	if response > 0 && response < 4 {
		game(response)
	}
}
