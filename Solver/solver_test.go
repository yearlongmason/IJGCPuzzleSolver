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

	// Actual value to be replaced by a loadPuzzle function
	actual := PuzzleState{3, 3,
		[][]string{
			{"0", "0", "0"},
			{"0", "0", "0"},
			{"0", "0", "0"}},
	}

	// If the puzzle states are not equal, then raise an error
	if !puzzleStatesEqual(expected, actual) {
		t.Errorf("Error!\nExpected: %v\nActual:   %v", expected, actual)
	}
}
