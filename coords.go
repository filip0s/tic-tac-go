package main

import "fmt"

// Coords stores coordinations of playing board cell in which one of the players
// try to put his marker
type Coords struct {
	row int
	col int
}

// Checks if the coordinates in struct are in the 0 to cellsPerSide range.
// Returns error in case, that one of the coordinates is outside the range, otherwise returns nil.
func (c Coords) checkCoordsDomain() error {

	if c.row >= cellsPerSide || c.col >= cellsPerSide || c.row < 0 || c.col < 0 {
		return fmt.Errorf("You are trying to access out of bounds cell.\nPlease use numbers between 0 and %d.\n",
			cellsPerSide-1)
	}

	return nil
}
