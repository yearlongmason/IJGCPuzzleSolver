package main

import (
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
	expected := "111|111|111|"

	// Get the actual value
	actual := turnRelicLeft(1, 1, puzzle).getHashableState()

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
	expected := ".11|11.|1.1|"

	// Get the actual value
	actual := turnRelicLeft(1, 1, puzzle3x3).getHashableState()

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
	expected := "110|110|000|"

	// Get the actual value
	actual := turnRelicLeft(0, 0, puzzle3x3).getHashableState()

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
	expected := "00100|00100|11111|00100|00100|"

	// Get the actual value
	actual := turnRelicRight(2, 2, puzzle5x5).getHashableState()

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
	expected := "00000|00.00|0.1.0|00.00|00000|"

	// Get the actual value
	actual := turnRelicRight(2, 2, puzzle5x5).getHashableState()

	// If the puzzle states are not equal, then raise an error
	if expected != actual {
		t.Errorf("Error!\nExpected: %v\nActual:   %v", expected, actual)
	}
}
