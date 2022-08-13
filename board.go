package main

import "fmt"

type Tile string

const cellsPerSide = 3

// Enum of Tile values
const (
	empty  Tile = " "
	circle Tile = "O"
	cross  Tile = "X"
)

type Board struct {
	matrix [cellsPerSide][cellsPerSide]Tile
}

func (b *Board) create() {
	// Initialize playing board with empty tiles
	for i := 0; i < cellsPerSide; i++ {
		for j := 0; j < cellsPerSide; j++ {
			b.matrix[i][j] = empty
		}
	}
}

func (b *Board) printRowDivider() {
	lengthOfLine := 4*cellsPerSide + 1

	for i := 0; i < lengthOfLine; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n")
}

func (b *Board) print() {

	b.printRowDivider()
	for i := 0; i < cellsPerSide; i++ {
		for j := 0; j < cellsPerSide; j++ {
			fmt.Printf("| %s ", b.matrix[i][j])
		}
		fmt.Printf("|\n")
		b.printRowDivider()
	}
}

func (b *Board) insert(position Coords, player Player) {
	b.matrix[position.row][position.col] = player.playerCharacter
}

//TODO: Check if cell where player is trying to insert is empty
