package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type PuzzleState struct {
	rows, columns int
	state         [][]string
}

func stringToPuzzleState(stringPuzzleState string) PuzzleState {
	// Create the variables needed to make an instance of PuzzleState
	slicePuzzleState := make([][]string, 0)
	numRows := len(strings.Split(stringPuzzleState, "\n"))
	numCols := len(strings.TrimSpace(strings.Split(stringPuzzleState, "\n")[0]))

	// Loop through each row in the puzzle input
	for row := range strings.SplitSeq(stringPuzzleState, "\n") {
		// Create a new row to keep track of
		row = strings.TrimSpace(row)
		currentRow := make([]string, 0)

		// Add each character to the new row
		for i := 0; i < len(row); i++ {
			currentRow = append(currentRow, string(row[i]))
		}

		// Add the new row to the puzzle state
		slicePuzzleState = append(slicePuzzleState, currentRow)
	}

	// Create and return a new PuzzleState
	newPuzzleState := PuzzleState{numRows, numCols, slicePuzzleState}
	return newPuzzleState
}

func loadPuzzle(fileName string) PuzzleState {
	// Load a puzzle from a file and return the puzzle state
	stringPuzzleState, err := os.ReadFile("../Puzzles/" + fileName)
	if err != nil {
		fmt.Printf("Error, could not read file: %s\n", fileName)
	}

	return stringToPuzzleState(string(stringPuzzleState))
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

// TODO -
// Create get successors function
// Create bfs to find shortest path

func main() {
	fmt.Println(loadPuzzle("GizehPuzzle.txt"))
}
