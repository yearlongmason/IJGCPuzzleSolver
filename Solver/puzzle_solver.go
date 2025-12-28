package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type PuzzleState struct {
	rows, columns, relicsUsed int
	state                     [][]string
	hashableState             string
}

func createNewPuzzleState(relicsUsed int, state [][]string) PuzzleState {
	// Create and return a new puzzle state given the relics used and the state of the puzzle
	// Create a string representation of the state from the state slice to be kept track of in a set
	var hashableState strings.Builder
	for _, row := range state {
		// Get the current row, and replace all instances of L and R with 1 to avoid
		// searching already explored paths
		currentRow := strings.Join(row, "")
		currentRow = strings.ReplaceAll(currentRow, "L", "1")
		currentRow = strings.ReplaceAll(currentRow, "R", "1")

		// Build on existing string
		hashableState.WriteString(currentRow)
		hashableState.WriteString("|") // Add separator between rows
	}

	return PuzzleState{
		rows:          len(state),
		columns:       len(state[0]),
		relicsUsed:    relicsUsed,
		state:         state,
		hashableState: hashableState.String(),
	}
}

func stringToPuzzleState(stringPuzzleState string) [][]string {
	// Take in a string representing a puzzle state and return the puzzle state as a slice
	slicePuzzleState := make([][]string, 0)

	// Loop through each row in the puzzle input
	for row := range strings.SplitSeq(stringPuzzleState, "\n") {
		// Remove whitespace from the row and add a new row to the slice
		row = strings.TrimSpace(row)
		slicePuzzleState = append(slicePuzzleState, strings.Split(row, ""))
	}

	return slicePuzzleState
}

func loadPuzzle(fileName string) PuzzleState {
	// Load a puzzle from a file and return the puzzle state
	stringPuzzleState, err := os.ReadFile("../Puzzles/" + fileName)
	if err != nil {
		fmt.Printf("Error, could not read file: %s\n", fileName)
	}

	// Get the puzzle state as a 2D slice of strings
	var state [][]string = stringToPuzzleState(string(stringPuzzleState))
	return createNewPuzzleState(0, state)
}

func puzzleStatesEqual(state1 PuzzleState, state2 PuzzleState) bool {
	// Return true if the puzzle states are equal, otherwise return false
	// Make sure all puzzle state attributes are equal
	rowsEqual := state1.rows == state2.rows
	columnsEqual := state1.columns == state2.columns
	relicsUsedEqual := state1.relicsUsed == state2.relicsUsed
	hashableStateEqual := state1.hashableState == state2.hashableState
	if !rowsEqual || !columnsEqual || !relicsUsedEqual || !hashableStateEqual {
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

func turnRelicLeft(row int, col int, currentState PuzzleState) PuzzleState {
	//updatedState := currentState.state

	return createNewPuzzleState(currentState.relicsUsed+1, currentState.state)
}
func turnRelicRight() {}
func getSuccessors()  {}

func main() {
	fmt.Println(loadPuzzle("3x3Puzzle.txt"))
}
