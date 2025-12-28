package main

import (
	"testing"
)

func TestLoadPuzzle3x3(t *testing.T) {
	// Hardcoded expected value
	expected := PuzzleState{3, 3, 0,
		[][]string{
			{"0", "0", "0"},
			{"0", "0", "0"},
			{"0", "0", "0"}},
		"000|000|000|",
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
	expected := PuzzleState{2, 2, 0,
		[][]string{
			{"0", "0"},
			{"0", "0"}},
		"00|00|",
	}

	// Get the actual puzzle from the 3x3Puzzle.txt test file
	actual := loadPuzzle("2x2Puzzle.txt")

	// If the puzzle states are not equal, then raise an error
	if !puzzleStatesEqual(expected, actual) {
		t.Errorf("Error!\nExpected: %v\nActual:   %v", expected, actual)
	}
}
