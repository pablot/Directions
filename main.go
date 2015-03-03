package main

import "fmt"

func menu() int {
	fmt.Println("Choose:")
	fmt.Println("1. Easy")
	fmt.Println("2. Medium")
	fmt.Println("3. Hard")
	fmt.Println("Other value. Quit")
	var response int
	fmt.Scanf("%d", &response)
	return response
}

func game(int mode) {

}

func main() {
	var response = menu()
	if response > 0 && response < 4 {
		game(response)
	}
}
