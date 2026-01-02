package main

import (
	"slices"
)

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

func getSuccessors(currentState PuzzleState) []PuzzleState {
	// Return a slice of all successors which is a new PuzzleState with
	// a relic in each slot turned left and right for each valid slot
	successors := make([]PuzzleState, 0)

	for rowIndex, row := range currentState.state {
		for colIndex, slotStatus := range row {
			// Check to make sure there is not already a relic in that spot
			// and that it is a valid slot
			slotHasRelic := slices.Contains([]string{"L", "R"}, slotStatus)
			if slotHasRelic || !isValidSlot(rowIndex, colIndex, currentState) {
				continue
			}

			// Append 2 new states:
			// One where a relic is inserted and turned to the left
			// One where a relic is inserted and turned to the right
			successors = append(successors, turnRelicLeft(rowIndex, colIndex, currentState))
			successors = append(successors, turnRelicRight(rowIndex, colIndex, currentState))
		}
	}

	return successors
}

func main() {
	puzzle5x5 := createNewPuzzleState(0, [][]string{
		{"0", "0", "0", "0", "0"},
		{"0", "0", "0", "0", "0"},
		{"0", "0", "0", "0", "0"},
		{"0", "0", "0", "0", "0"},
		{"0", "0", "0", "0", "0"}})
	puzzle5x5 = turnRelicRight(2, 2, puzzle5x5)
	puzzle5x5 = turnRelicLeft(1, 1, puzzle5x5)
	puzzle5x5.printPuzzleState()
}
