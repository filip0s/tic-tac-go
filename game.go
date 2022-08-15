package main

import (
	"fmt"
	"os"
)

type Game struct {
	playingBoard  Board
	player1       Player
	player2       Player
	currentPlayer Player
	isOver        bool
}

func (g *Game) create() {
	g.playingBoard.create()
	g.player1.create(cross)
	g.player2.create(circle)
	g.currentPlayer = g.player1
	g.isOver = false
}

func (g *Game) switchPlayers() {
	if g.currentPlayer == g.player1 {
		g.currentPlayer = g.player2
	} else {
		g.currentPlayer = g.player1
	}

}

// Collects input from one of players, checks it and tries to put Player's
// character on the Board based on the input
func (g *Game) handleInput() {
	var input Coords

	fmt.Printf("[Player %s] enter row and column (0 - %d): ", g.currentPlayer.playerCharacter, cellsPerSide-1)

	_, err := fmt.Scanf("%d %d \n", &input.row, &input.col)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Please enter valid number of row and column.\n") // Rewrite this with the use of errors.New()
		g.handleInput()
	}

	err = input.checkCoordsDomain()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s", err)
		g.handleInput()
	}

	g.playingBoard.insert(input, g.currentPlayer)
}

func (g *Game) play() {
	const maximumRounds = cellsPerSide * cellsPerSide
	var roundCounter = 0

	for roundCounter < maximumRounds && !g.isOver {
		// TODO: Implement game logic
		g.playingBoard.print()

		g.handleInput()
		g.switchPlayers()
		roundCounter++
	}
}
