package main

import "directions/game"

func main() {
	game.Menu()
	var response = game.GetResponse()
	if response > 0 && response < 4 {
		game.Game(response)
	}
}
