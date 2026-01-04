package main

import (
	"container/heap"
	"slices"
	"testing"
)

func TestLoadPuzzle3x3(t *testing.T) {
	// Hardcoded expected value
	expected := PuzzleState{0,
		[][]string{
			{"0", "0", "0"},
			{"0", "0", "0"},
			{"0", "0", "0"}},
	}

	// Get the actual puzzle from the 3x3Puzzle.txt test file
	actual := loadPuzzle("3x3Puzzle.txt")

	// If the puzzle states are not equal, then raise an error
	if !puzzleStatesEqual(expected, actual) {
		t.Errorf("Error!\nExpected: %v\nActual:   %v", expected, actual)
	}
}

func TestLoadPuzzle2x2(t *testing.T) {
	// Hardcoded expected value
	expected := PuzzleState{0,
		[][]string{
			{"0", "0"},
			{"0", "0"}},
	}

	// Get the actual puzzle from the 3x3Puzzle.txt test file
	actual := loadPuzzle("2x2Puzzle.txt")

	// If the puzzle states are not equal, then raise an error
	if !puzzleStatesEqual(expected, actual) {
		t.Errorf("Error!\nExpected: %v\nActual:   %v", expected, actual)
	}
}

func TestTurnRelicLeft3x3Middle(t *testing.T) {
	puzzle := PuzzleState{0,
		[][]string{
			{"0", "0", "0"},
			{"0", "0", "0"},
			{"0", "0", "0"}},
	}

	// Hardcoded expected value
	expected := "111|1L1|111|"

	// Get the actual value
	actual := turnRelicLeft(1, 1, puzzle).getStateString()

	// If the puzzle states are not equal, then raise an error
	if expected != actual {
		t.Errorf("Error!\nExpected: %v\nActual:   %v", expected, actual)
	}
}

func TestTurnRelicLeft3x3MiddleBlocked(t *testing.T) {
	puzzle3x3 := PuzzleState{0,
		[][]string{
			{".", "0", "0"},
			{"0", "0", "."},
			{"0", ".", "0"}},
	}

	// Hardcoded expected value
	expected := ".11|1L.|1.1|"

	// Get the actual value
	actual := turnRelicLeft(1, 1, puzzle3x3).getStateString()

	// If the puzzle states are not equal, then raise an error
	if expected != actual {
		t.Errorf("Error!\nExpected: %v\nActual:   %v", expected, actual)
	}
}

func TestTurnRelicLeft3x3TopLeft(t *testing.T) {
	puzzle3x3 := PuzzleState{0,
		[][]string{
			{"0", "0", "0"},
			{"0", "0", "0"},
			{"0", "0", "0"}},
	}

	// Hardcoded expected value
	expected := "L10|110|000|"

	// Get the actual value
	actual := turnRelicLeft(0, 0, puzzle3x3).getStateString()

	// If the puzzle states are not equal, then raise an error
	if expected != actual {
		t.Errorf("Error!\nExpected: %v\nActual:   %v", expected, actual)
	}
}

func TestTurnRelicRight5x5Middle(t *testing.T) {
	puzzle5x5 := createNewPuzzleState(0, [][]string{
		{"0", "0", "0", "0", "0"},
		{"0", "0", "0", "0", "0"},
		{"0", "0", "0", "0", "0"},
		{"0", "0", "0", "0", "0"},
		{"0", "0", "0", "0", "0"}})

	// Hardcoded expected value
	expected := "00100|00100|11R11|00100|00100|"

	// Get the actual value
	actual := turnRelicRight(2, 2, puzzle5x5).getStateString()

	// If the puzzle states are not equal, then raise an error
	if expected != actual {
		t.Errorf("Error!\nExpected: %v\nActual:   %v", expected, actual)
	}
}

func TestTurnRelicRight5x5MiddleBlocked(t *testing.T) {
	puzzle5x5 := createNewPuzzleState(0, [][]string{
		{"0", "0", "0", "0", "0"},
		{"0", "0", ".", "0", "0"},
		{"0", ".", "0", ".", "0"},
		{"0", "0", ".", "0", "0"},
		{"0", "0", "0", "0", "0"}})

	// Hardcoded expected value
	expected := "00000|00.00|0.R.0|00.00|00000|"

	// Get the actual value
	actual := turnRelicRight(2, 2, puzzle5x5).getStateString()

	// If the puzzle states are not equal, then raise an error
	if expected != actual {
		t.Errorf("Error!\nExpected: %v\nActual:   %v", expected, actual)
	}
}

func TestGetSuccessors2x2(t *testing.T) {
	puzzle2x2 := createNewPuzzleState(0, [][]string{
		{"0", "0"},
		{"0", "0"}})

	// Hardcoded expected value
	expected := []string{"R1|10|", "1R|01|", "10|R1|", "01|1R|",
		"L1|11|", "1L|11|", "11|L1|", "11|1L|"}

	// Get actual values
	actual := make([]string, 0)
	for _, successor := range getSuccessors(puzzle2x2) {
		actual = append(actual, successor.getStateString())
	}

	// Sort slices and check if they are equal
	slices.Sort(expected)
	slices.Sort(actual)
	if !slices.Equal(expected, actual) {
		t.Errorf("Error!\nExpected: %v\nActual:   %v", expected, actual)
	}
}

func TestGetSuccessors2x2WithOneFilled(t *testing.T) {
	puzzle2x2 := createNewPuzzleState(0, [][]string{
		{"1", "0"},
		{"0", "0"}})

	// Hardcoded expected value
	expected := []string{"R1|10|", "1R|01|", "10|R1|", "11|1R|",
		"L1|11|", "1L|11|", "11|L1|", "11|1L|"}

	// Get actual values
	actual := make([]string, 0)
	for _, successor := range getSuccessors(puzzle2x2) {
		actual = append(actual, successor.getStateString())
	}

	// Sort slices and check if they are equal
	slices.Sort(expected)
	slices.Sort(actual)
	if !slices.Equal(expected, actual) {
		t.Errorf("Error!\nExpected: %v\nActual:   %v", expected, actual)
	}
}

func TestNoSuccessors(t *testing.T) {
	puzzle2x2 := createNewPuzzleState(0, [][]string{
		{"L", "L"},
		{"L", "L"}})

	// Hardcoded expected value
	expected := make([]string, 0)

	// Get actual values
	actual := make([]string, 0)
	for _, successor := range getSuccessors(puzzle2x2) {
		actual = append(actual, successor.getStateString())
	}

	// Check that both are empty
	if !slices.Equal(expected, actual) {
		t.Errorf("Error!\nExpected: %v\nActual:   %v", expected, actual)
	}
}

func TestPuzzleStatePriorityQueue(t *testing.T) {
	// Create priority queue and fill with
	priorityQueue := &PuzzleStatePriorityQueue{}
	heap.Push(priorityQueue, createNewPuzzleState(0, [][]string{{"0", "0"}, {"0", "0"}}))
	heap.Push(priorityQueue, createNewPuzzleState(0, [][]string{{"1", "0"}, {"0", "0"}}))
	heap.Push(priorityQueue, createNewPuzzleState(0, [][]string{{"1", "1"}, {"0", "0"}}))
	heap.Push(priorityQueue, createNewPuzzleState(0, [][]string{{"R", "1"}, {"1", "0"}}))
	heap.Push(priorityQueue, createNewPuzzleState(0, [][]string{{"L", "1"}, {"1", "1"}}))

	// Expect to get the puzzle state with the most
	expected := "L1|11|"
	actual := heap.Pop(priorityQueue).(PuzzleState).getStateString()

	if expected != actual {
		t.Errorf("Error!\nExpected: %v\nActual:   %v", expected, actual)
	}
}
