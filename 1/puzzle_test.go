package main

import (
	"testing"
)

var testInput string = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestPuzzleOne(t *testing.T) {
	result, _ := puzzleOne(testInput)
	if result != 11 {
		t.Errorf("Result was not correct. Got: %d Want: %d", result, 11)
	}
}

func TestPuzzleTwo(t *testing.T) {
	result, _ := puzzleTwo(testInput)
	if result != 31 {
		t.Errorf("Result was not correct. Got: %d Want: %d", result, 31)
	}
}
