package game

import "fmt"

//Menu - Shows Menu for game
func Menu() {
	fmt.Println("Choose:")
	fmt.Println("1. Easy")
	fmt.Println("2. Medium")
	fmt.Println("3. Hard")
	fmt.Println("Other value. Quit")
}

//GetResponse -  form user
func GetResponse() int {
	var response int
	fmt.Scanf("%d", &response)
	return response
}

//Game - The game itself
func Game(mode int) {
}
