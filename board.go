package main

import (
	"errors"
	"fmt"
)

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

func (b *Board) insert(pos Coords, player Player) {
	b.matrix[pos.row][pos.col] = player.playerCharacter
}

// Checks validity of coordinates. Method tests if coordinates are on the playing
// board and if the cell is empty.
func (b *Board) check(pos Coords) error {
	err := checkCoordsDomain(pos)
	if err != nil {
		return err
	}

	err = b.checkEmpty(pos)
	if err != nil {
		return err
	}

	return nil
}

// Checks if the coordinates in struct are in the 0 to cellsPerSide-1 range.
// Returns error in case, that one of the coordinates is outside the range, otherwise returns nil.
func checkCoordsDomain(pos Coords) error {
	var maximumPosition = cellsPerSide - 1

	if pos.row > maximumPosition || pos.col > maximumPosition || pos.row < 0 || pos.col < 0 {
		errOutput := fmt.Sprintf("You are trying to access out of bounds cell.\nPlease use numbers between 0 and %d.",
			maximumPosition)
		return errors.New(errOutput)
	}
	return nil
}

// Checks if coordinates passed as parameter are empty.
func (b *Board) checkEmpty(pos Coords) error {
	if b.matrix[pos.row][pos.col] != empty {
		return errors.New("inserting into full cell")
	}

	return nil
}
