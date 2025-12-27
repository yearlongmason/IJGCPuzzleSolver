package main

import (
	"testing"
)

func TestLoadPuzzle(t *testing.T) {
	// Hardcoded expected value
	expected := PuzzleState{3, 3,
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
