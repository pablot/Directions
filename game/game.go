package game

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type generator func(int) string
type response func() string
type counting func(int)

//Menu - Shows Menu for game
func Menu() int {
	fmt.Println("Choose:")
	fmt.Println("1. Easy")
	fmt.Println("2. Medium")
	fmt.Println("3. Hard")
	fmt.Println("Other value. Quit")
	return getResponseMenu()
}

//Game - The game itself
func Game(mode int) {
	if game(mode*3, generateDirections, getResponseGame, countingToZero) {
		fmt.Println("Gratualtion You won!")
	} else {
		fmt.Println("Not this time!")
	}
}

func getResponseMenu() int {
	var response int
	fmt.Scanf("%d\n", &response)
	return response
}

func game(quantity int, genDirections generator, getRspGame response, countingFromX counting) bool {
	clearScreen()

	fmt.Println("Directions:")
	fromComputer := genDirections(quantity)
	fmt.Println(fromComputer)
	fmt.Println()
	fmt.Println("Remember this directions!")

	countingFromX(3)

	clearScreen()

	fmt.Println("Now write directions as <>^v:")
	fromUser := getRspGame()

	return compareResults(fromUser, fromComputer)

}

func compareResults(fromUser string, fromComputer string) bool {
	return fromUser == fromComputer
}

func getResponseGame() string {
	var response string
	fmt.Scanf("%s\n", &response)
	return response
}

func countingToZero(from int) {
	for i := from; i > 0; i-- {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(i)
	}
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func generateDirections(quantity int) string {
	var result string
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < quantity; i++ {
		value, err := generateDirection(rand.Intn(3))
		if err == nil {
			result += value
		}
	}
	return result
}

func generateDirection(from int) (string, error) {
	switch from {
	case 0:
		return "<", nil
	case 1:
		return ">", nil
	case 2:
		return "^", nil
	case 3:
		return "v", nil
	default:
		return "?", errors.New("Argument from should be in 0 - 3 range!")
	}
}
