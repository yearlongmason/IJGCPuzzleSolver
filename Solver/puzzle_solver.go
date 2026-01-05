package main

import (
	"container/heap"
	"fmt"
	"slices"
)

type PuzzleStatePriorityQueue []PuzzleState

// All functions required to implement heap.Interface
func (q PuzzleStatePriorityQueue) Len() int      { return len(q) }
func (q PuzzleStatePriorityQueue) Swap(i, j int) { q[i], q[j] = q[j], q[i] }
func (q PuzzleStatePriorityQueue) Less(i, j int) bool {
	return getPuzzleStateHeuristic(q[i]) < getPuzzleStateHeuristic(q[j])
}
func (q *PuzzleStatePriorityQueue) Push(newPuzzleState any) {
	*q = append(*q, newPuzzleState.(PuzzleState))
}
func (q *PuzzleStatePriorityQueue) Pop() any {
	// Pop a puzzle state off the end, and return it
	n := len(*q)
	oldPriorityQueue := *q
	poppedPuzzleState := oldPriorityQueue[n-1]
	*q = oldPriorityQueue[:n-1]
	return poppedPuzzleState
}

func getPuzzleStateHeuristic(puzzleState PuzzleState) int {
	// Add 1 for each activated slot. Return heuristic (int)
	heuristic := 0

	activatedSlots := []string{"1", "L", "R"}
	for _, row := range puzzleState.state {
		for _, character := range row {
			if slices.Contains(activatedSlots, character) {
				heuristic += 1
			}
		}
	}

	// Define some weight to incentivise using less relics
	// If this weight is too high it will eventually devolve back into BFS
	relicUsedWeight := 5

	// Multiply by negative 1 because we want to maximize the number of activated slots
	return (heuristic * -1) + (puzzleState.relicsUsed * relicUsedWeight)
}

func turnRelicLeft(row int, col int, currentState PuzzleState) PuzzleState {
	// Activate all 8 valid slots surrounding the current slot and return a new PuzzleState
	// Create new updated state and set the current cell to "L" to signify a relic turned left
	newPuzzleState := currentState.copy()
	updatedState := newPuzzleState.state
	updatedState[row][col] = "L"
	newPuzzleState.relicsUsed += 1
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
	return newPuzzleState
}

func turnRelicRight(row int, col int, currentState PuzzleState) PuzzleState {
	// Activate all slots on the same row and column until an invalid slot is hit
	// Return a new PuzzleState
	newPuzzleState := currentState.copy()
	updatedState := newPuzzleState.state
	updatedState[row][col] = "R"
	newPuzzleState.relicsUsed += 1

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

	return newPuzzleState
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

func findSolutionBFS(startPuzzleState PuzzleState) PuzzleState {
	// Use bfs to find the solution with the least possible relics
	queue := []PuzzleState{startPuzzleState}
	explored := map[string]bool{startPuzzleState.getSlotStatusesString(): true}

	// Continue looping until there is nothing left in the queue
	for len(queue) != 0 {
		currentPuzzleState := queue[0]
		queue = queue[1:]

		for _, successor := range getSuccessors(currentPuzzleState) {
			// Check if the successor is the solved state
			if isSolved(successor) {
				return successor
			}

			// Only if we have not already seen a state with this formation of activated slots
			if !explored[successor.getSlotStatusesString()] {
				// Mark as explored and add the PuzzleState to the queue
				explored[successor.getSlotStatusesString()] = true
				queue = append(queue, successor)
			}
		}
	}

	// This will never run since it is impossible to make an unsolvable ancient relic puzzle
	return startPuzzleState
}

func findSolutionAStar(startPuzzleState PuzzleState) PuzzleState {
	// Use A* to find a solution (does not return solution with least possible relics)
	priorityQueue := &PuzzleStatePriorityQueue{}
	heap.Push(priorityQueue, startPuzzleState)
	explored := map[string]bool{startPuzzleState.getSlotStatusesString(): true}

	// Continue looping until there is nothing left in the queue
	for priorityQueue.Len() != 0 {
		currentPuzzleState := heap.Pop(priorityQueue)

		for _, successor := range getSuccessors(currentPuzzleState.(PuzzleState)) {
			// Check if the successor is the solved state
			if isSolved(successor) {
				return successor
			}

			// Only if we have not already seen a state with this formation of activated slots
			if !explored[successor.getSlotStatusesString()] {
				// Mark as explored and add the PuzzleState to the queue
				explored[successor.getSlotStatusesString()] = true
				heap.Push(priorityQueue, successor)
			}
		}
	}

	// This will never run since it is impossible to make an unsolvable ancient relic puzzle
	return startPuzzleState
}

func main() {
	puzzle := loadPuzzle("GizehPuzzle.txt")

	// fmt.Println("BFS:")
	// puzzleBFS := findSolutionBFS(puzzle)
	// puzzleBFS.printPuzzleState()
	// fmt.Println(puzzleBFS.relicsUsed)

	fmt.Println("\nA*:")
	puzzleAStar := findSolutionAStar(puzzle)
	puzzleAStar.printPuzzleState()
	fmt.Println(puzzleAStar.relicsUsed)
}
