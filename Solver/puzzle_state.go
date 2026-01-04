package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type PuzzleState struct {
	relicsUsed int
	state      [][]string
}

func createNewPuzzleState(relicsUsed int, state [][]string) PuzzleState {
	// Create and return a new PuzzleState
	return PuzzleState{
		relicsUsed: relicsUsed,
		state:      state,
	}
}

func (puzzleState PuzzleState) getNumRows() int {
	// Getter for the number of rows in a puzzle
	return len(puzzleState.state)
}

func (puzzleState PuzzleState) getNumColumns() int {
	// Getter for finding the number of columns in a puzzle state
	// This assumes that all columns are the same number of rows which they should be
	return len(puzzleState.state[0])
}

func (puzzleState PuzzleState) getSlotStatusesString() string {
	// Create a string representation of the slot statuses from the state slice
	// We only want the slot status so replace relics with "1"
	var slotStatusesString strings.Builder
	for _, row := range puzzleState.state {
		// Get the current row, and replace all instances of "L" and "R" with "1"
		currentRow := strings.Join(row, "")
		currentRow = strings.ReplaceAll(currentRow, "L", "1")
		currentRow = strings.ReplaceAll(currentRow, "R", "1")

		// Build on existing string
		slotStatusesString.WriteString(currentRow)
		slotStatusesString.WriteString("|") // Add separator between rows
	}

	return slotStatusesString.String()
}

func (puzzleState PuzzleState) getStateString() string {
	// Create a string representation of the state from the state slice
	// We want to know where the relics are too, so keep in "L" and "R"
	var stateString strings.Builder
	for _, row := range puzzleState.state {
		// Build on existing string
		stateString.WriteString(strings.Join(row, ""))
		stateString.WriteString("|") // Add separator between rows
	}

	return stateString.String()
}

func (puzzleState PuzzleState) copy() PuzzleState {
	// Return a copy of the PuzzleState
	copiedState := make([][]string, 0)

	// "Deepcopy" the puzzle state to avoid copying a reference to the array
	for rowIndex, row := range puzzleState.state {
		copiedState = append(copiedState, make([]string, 0))
		for _, character := range row {
			copiedState[rowIndex] = append(copiedState[rowIndex], character)
		}
	}

	return createNewPuzzleState(puzzleState.relicsUsed, copiedState)
}

func (puzzleState PuzzleState) printPuzzleState() {
	// Nicely print out the current PuzzleState
	for _, row := range puzzleState.state {
		fmt.Println(strings.Join(row, ""))
	}
}

func puzzleStatesEqual(state1 PuzzleState, state2 PuzzleState) bool {
	// Return true if the puzzle states are equal, otherwise return false
	relicsUsedEqual := state1.relicsUsed == state2.relicsUsed
	if !relicsUsedEqual {
		return false
	}

	// Loop through each row in both puzzle states and check if they are equal
	// If they are not equal then return false
	for i := 0; i < state1.getNumRows(); i++ {
		if !slices.Equal(state1.state[i], state2.state[i]) {
			return false
		}
	}
	return true
}

func isValidSlot(row int, col int, currentState PuzzleState) bool {
	// Returns true if the given row and column is a valid slot (exists and is not ".")
	// Check for invalid cases: row or column are out of bounds or the cell is not an actual slot
	isValidRow := row >= 0 && row < currentState.getNumRows()
	isValidCol := col >= 0 && col < currentState.getNumColumns()
	if !isValidRow || !isValidCol || currentState.state[row][col] == "." {
		return false
	}

	return true
}

func loadPuzzle(fileName string) PuzzleState {
	// Load a puzzle from a file and return the puzzle state
	stringPuzzleState, err := os.ReadFile("../Puzzles/" + fileName)
	if err != nil {
		fmt.Printf("Error, could not read file: %s\n", fileName)
	}

	// Get the puzzle state as a 2D slice of strings
	slicePuzzleState := make([][]string, 0)

	// Loop through each row in the puzzle input and add to the slice puzzle state
	for row := range strings.SplitSeq(string(stringPuzzleState), "\n") {
		row = strings.TrimSpace(row)
		slicePuzzleState = append(slicePuzzleState, strings.Split(row, ""))
	}

	return createNewPuzzleState(0, slicePuzzleState)
}
