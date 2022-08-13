package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	var game Game
	game.create()
	game.playingBoard.print()
}
