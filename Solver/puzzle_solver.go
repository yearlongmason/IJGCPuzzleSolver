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

func (puzzleState PuzzleState) getNumRows() int {
	// Getter for the number of rows in a puzzle
	return len(puzzleState.state)
}

func (puzzleState PuzzleState) getNumColumns() int {
	// Getter for finding the number of columns in a puzzle state
	// This assumes that all columns are the same number of rows which they should be
	return len(puzzleState.state[0])
}

func (puzzleState PuzzleState) getHashableState() string {
	// Getter for a hashable state
	// Create a string representation of the state from the state slice to be kept track of in a set
	// We only want to keep track of active and inactive slots, so replace relics with "1"
	var hashableState strings.Builder
	for _, row := range puzzleState.state {
		// Get the current row, and replace all instances of "L" and "R" with "1"
		currentRow := strings.Join(row, "")
		currentRow = strings.ReplaceAll(currentRow, "L", "1")
		currentRow = strings.ReplaceAll(currentRow, "R", "1")

		// Build on existing string
		hashableState.WriteString(currentRow)
		hashableState.WriteString("|") // Add separator between rows
	}

	return hashableState.String()
}

func (puzzleState PuzzleState) printPuzzleState() {
	// Getter for the number of rows in a puzzle
	for _, row := range puzzleState.state {
		fmt.Println(strings.Join(row, ""))
	}
}

func createNewPuzzleState(relicsUsed int, state [][]string) PuzzleState {
	// Create and return a new puzzle state given the relics used and the state of the puzzle
	return PuzzleState{
		relicsUsed: relicsUsed,
		state:      state,
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
	rowsEqual := state1.getNumRows() == state2.getNumRows()
	columnsEqual := state1.getNumColumns() == state2.getNumColumns()
	relicsUsedEqual := state1.relicsUsed == state2.relicsUsed
	hashableStateEqual := state1.getHashableState() == state2.getHashableState()
	if !rowsEqual || !columnsEqual || !relicsUsedEqual || !hashableStateEqual {
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

func turnRelicLeft(row int, col int, currentState PuzzleState) PuzzleState {
	// Activate all 8 valid slots surrounding the current slot and return a new PuzzleState
	// Create new updated state and set the current cell to "L" to signify a relic turned left
	updatedState := currentState.state
	updatedState[row][col] = "L"
	coordChanges := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	// Change all the valid points surrounding the relic
	for _, change := range coordChanges {
		// If the shifted position is valid then set the slot to activated in the updated state
		newRow := row + change[0]
		newCol := col + change[1]
		if isValidSlot(newRow, newCol, currentState) {
			// If the current slot is covered by a relic just skip over it
			isAlreadyActive := slices.Contains([]string{"L", "R", "1"}, updatedState[newRow][newCol])
			if isAlreadyActive {
				continue
			}
			updatedState[newRow][newCol] = "1"
		}
	}

	// Return a new PuzzleState with the updated number of relics used and the new state
	return createNewPuzzleState(currentState.relicsUsed+1, updatedState)
}

func turnRelicRight(row int, col int, currentState PuzzleState) PuzzleState {
	// Activate all slots on the same row and column until an invalid slot is hit
	// Return a new PuzzleState
	updatedState := currentState.state
	updatedState[row][col] = "R"

	// Mark all slots on the row as activated until an invalid slot is hit
	// Look at entire row ahead
	for currentCol := col + 1; currentCol < currentState.getNumColumns(); currentCol++ {
		if isValidSlot(row, currentCol, currentState) {
			// If the current slot is covered by a relic just skip over it
			isAlreadyActive := slices.Contains([]string{"L", "R", "1"}, updatedState[row][currentCol])
			if isAlreadyActive {
				continue
			}
			updatedState[row][currentCol] = "1"
		} else {
			break
		}
	}
	// Look at entire row behind
	for currentCol := col - 1; currentCol >= 0; currentCol-- {
		if isValidSlot(row, currentCol, currentState) {
			// If the current slot is covered by a relic just skip over it
			isAlreadyActive := slices.Contains([]string{"L", "R", "1"}, updatedState[row][currentCol])
			if isAlreadyActive {
				continue
			}
			updatedState[row][currentCol] = "1"
		} else {
			break
		}
	}

	// Mark all slots on the column as activated until an invalid slot is hit
	// Look at entire column below
	for currentRow := row + 1; currentRow < currentState.getNumRows(); currentRow++ {
		if isValidSlot(currentRow, col, currentState) {
			// If the current slot is covered by a relic just skip over it
			isAlreadyActive := slices.Contains([]string{"L", "R", "1"}, updatedState[currentRow][col])
			if isAlreadyActive {
				continue
			}
			updatedState[currentRow][col] = "1"
		} else {
			break
		}
	}
	// Look at entire column above
	for currentRow := row - 1; currentRow >= 0; currentRow-- {
		if isValidSlot(currentRow, col, currentState) {
			// If the current slot is covered by a relic just skip over it
			isAlreadyActive := slices.Contains([]string{"L", "R", "1"}, updatedState[currentRow][col])
			if isAlreadyActive {
				continue
			}
			updatedState[currentRow][col] = "1"
		} else {
			break
		}
	}

	return createNewPuzzleState(currentState.relicsUsed+1, updatedState)
}

func getSuccessors() {}

func main() {
	//fmt.Println(loadPuzzle("3x3Puzzle.txt"))
	//fmt.Println(turnRelicLeft(1, 1, loadPuzzle("GizehPuzzle.txt")))
	puzzle5x5 := createNewPuzzleState(0, [][]string{
		{"0", "0", "0", "0", "0"},
		{"0", "0", "0", "0", "0"},
		{"0", "0", "0", "0", "0"},
		{"0", "0", "0", "0", "0"},
		{"0", "0", "0", "0", "0"}})
	puzzle5x5 = turnRelicRight(2, 2, puzzle5x5)
	//puzzle5x5 = turnRelicLeft(1, 1, puzzle5x5)
	puzzle5x5.printPuzzleState()
	fmt.Println(puzzle5x5.getHashableState())
}
