package main

import (
	"fmt"
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
	var input Position

	fmt.Printf("[Player %s] enter ROW and COLUMN (0 - %d): ", g.currentPlayer.playerCharacter, cellsPerSide-1)

	_, err := fmt.Scan(&input.row, &input.col)

	if err != nil {
		fmt.Println(err)
		g.handleInput()
		return
	}

	err = g.playingBoard.check(input)
	if err != nil {
		fmt.Println(err)
		g.handleInput()
		return
	}
	g.playingBoard.insert(input, g.currentPlayer)
	g.lastPosition = input
}

func adjustStartPos(cell int) int {
	if cell < 0 {
		return 0
	}
	return cell
}

func adjustEndPos(cell int) int {
	maxPosition := cellsPerSide - 1
	if cell >= maxPosition {
		return maxPosition
	}
	return cell
}

// Checks if there are 3 characters in row, where was the latest insertion
func (g *Game) checkHorizontal() bool {
	var (
		col      = g.lastPosition.col
		startRow = adjustStartPos(g.lastPosition.row - 1)
		endRow   = adjustEndPos(g.lastPosition.row + 1)
		isWin    = true
	)

	for row := startRow; row <= endRow; row++ {
		isWin = isWin && (g.playingBoard.matrix[row][col] == g.currentPlayer.playerCharacter)
	}

	return isWin
}

// Check if there are 3 same characters in vertical orientation (column) around the latest insertion
func (g *Game) checkVertical() bool {
	var (
		row      = g.lastPosition.row
		startCol = adjustStartPos(g.lastPosition.col - 1)
		endCol   = adjustEndPos(g.lastPosition.col + 1)
		isWin    = true
	)

	for col := startCol; col <= endCol; col++ {
		isWin = isWin && (g.playingBoard.matrix[row][col] == g.currentPlayer.playerCharacter)
	}

	return isWin
}

func (g *Game) checkMainDiagonal() bool {
	return false
}

func (g *Game) checkSideDiagonal() bool {
	return false
}

func (g *Game) checkWin() bool {
	return g.checkHorizontal() || g.checkVertical() || g.checkMainDiagonal() || g.checkSideDiagonal()
}

func (g *Game) play() {
	const maximumRounds = cellsPerSide * cellsPerSide
	var roundCounter = 0

	for roundCounter < maximumRounds && !g.isOver {
		// TODO: Implement game logic
		g.playingBoard.print()

		g.handleInput()
		g.switchPlayers()

		//TODO: Implement check for win
		roundCounter++
	}
}
