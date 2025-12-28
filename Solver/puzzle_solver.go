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
	// Create the hashable state from the
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
	// Create the variables needed to make an instance of PuzzleState
	slicePuzzleState := make([][]string, 0)

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

	return slicePuzzleState
}

func loadPuzzle(fileName string) PuzzleState {
	// Load a puzzle from a file and return the puzzle state
	stringPuzzleState, err := os.ReadFile("../Puzzles/" + fileName)
	if err != nil {
		fmt.Printf("Error, could not read file: %s\n", fileName)
	}

	state := stringToPuzzleState(string(stringPuzzleState))
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

// TODO -
// Create get successors function
// Create bfs to find shortest path
func turnRelicLeft()  {}
func turnRelicRight() {}
func getSuccessors()  {}

func main() {
	fmt.Println(loadPuzzle("3x3Puzzle.txt"))
}
