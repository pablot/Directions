package main

import "directions/game"

func main() {
	var response = game.Menu()
	if response > 0 && response < 4 {
		game.Game(response)
	}
}
