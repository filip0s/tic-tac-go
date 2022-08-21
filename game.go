package main

import (
	"fmt"
)

type Game struct {
	playingBoard  Board
	player1       Player
	player2       Player
	currentPlayer Player
	state         State
	lastPosition  Position
}

type State int

const (
	inProgress State = iota
	player1Win State = iota
	player2Win State = iota
	tie        State = iota
)

func (g *Game) create() {
	g.playingBoard.create()
	g.player1.create(cross)
	g.player2.create(circle)
	g.currentPlayer = g.player1
	g.state = inProgress
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

// Performs a check if there was certain number (`goal` value) of current player's
// character in row inside the provided slice. Returns boolean value depending on
// if the goal was reached.
func (g *Game) checkBoardSlice(checkedSlice []Tile) bool {
	var (
		goal          = 3
		isGoalReached = false
		counter       = 0
	)

	for _, val := range checkedSlice {
		if val == g.currentPlayer.playerCharacter {
			counter++
			if counter == goal {
				isGoalReached = true
				break
			}
		} else {
			counter = 0
		}

	}

	return isGoalReached
}

// Checks if there are 3 characters in row, where was the latest insertion
func (g *Game) checkHorizontal() bool {
	var (
		row         = g.lastPosition.row
		startCol    = adjustStartPos(g.lastPosition.col - 2)
		endCol      = adjustEndPos(g.lastPosition.col + 2)
		checkedPart []Tile
	)

	for col := startCol; col <= endCol; col++ {
		checkedPart = append(checkedPart, g.playingBoard.matrix[row][col])
	}

	return g.checkBoardSlice(checkedPart)
}

// Checks if there are 3 same characters in vertical orientation (column) around the latest insertion
func (g *Game) checkVertical() bool {
	var (
		col         = g.lastPosition.col
		startRow    = adjustStartPos(g.lastPosition.row - 2)
		endRow      = adjustEndPos(g.lastPosition.row + 2)
		checkedPart []Tile
	)

	for row := startRow; row <= endRow; row++ {
		checkedPart = append(checkedPart, g.playingBoard.matrix[row][col])
	}

	return g.checkBoardSlice(checkedPart)
}

func (g *Game) checkMainDiagonal() bool {
	var (
		startRow    = adjustStartPos(g.lastPosition.row - 2)
		startCol    = adjustStartPos(g.lastPosition.col - 2)
		endRow      = adjustEndPos(g.lastPosition.row + 2)
		endCol      = adjustEndPos(g.lastPosition.col + 2)
		checkedPart []Tile
	)

	row := startRow
	col := startCol
	for row <= endRow && col <= endCol {
		checkedPart = append(checkedPart, g.playingBoard.matrix[row][col])
		row++
		col++
	}

	return g.checkBoardSlice(checkedPart)
}

func (g *Game) checkAntiDiagonal() bool {
	var (
		startRow    = adjustStartPos(g.lastPosition.row - 2)
		startCol    = adjustEndPos(g.lastPosition.col + 2)
		endRow      = adjustEndPos(g.lastPosition.row + 2)
		endCol      = adjustStartPos(g.lastPosition.col - 2)
		checkedPart []Tile
	)

	row := startRow
	col := startCol
	for row <= endRow && col >= endCol {
		checkedPart = append(checkedPart, g.playingBoard.matrix[row][col])
		row++
		col--
	}

	return g.checkBoardSlice(checkedPart)
}

// Checks if there was met any of win conditions after latest insertion.
func (g *Game) checkWin() bool {
	return g.checkHorizontal() || g.checkVertical() || g.checkMainDiagonal() || g.checkAntiDiagonal()
}

func (g *Game) handleWinState() {
	if g.currentPlayer == g.player1 {
		g.state = player1Win
		return
	}

	g.state = player2Win
}

func (g *Game) play() {
	const maximumRounds = cellsPerSide * cellsPerSide
	var roundCounter = 0

	for roundCounter < maximumRounds && g.state == inProgress {
		g.playingBoard.print()

		g.handleInput()
		if g.checkWin() {
			g.handleWinState()
			return
		}

		g.switchPlayers()
		roundCounter++
	}

	g.state = tie
}
