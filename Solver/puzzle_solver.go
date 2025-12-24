package main

import (
	"fmt"
	"slices"
)

type PuzzleState struct {
	rows, columns int
	state         [][]string
}

func puzzleStatesEqual(state1 PuzzleState, state2 PuzzleState) bool {
	// Return true if the puzzle states are equal, otherwise return false
	// Basic check to make sure the puzzle states are the same size
	if state1.rows != state2.rows || state1.columns != state2.columns {
		return false
	}

	// Loop through each row in both puzzle states and check if they are equal
	// If they are not equal then return false
	for i := 0; i < state1.rows; i++ {
		if !slices.Equal(state1.state[i], state2.state[i]) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(!slices.Equal([]string{"1", "0", "3"}, []string{"1", "2", "3"}))
}
