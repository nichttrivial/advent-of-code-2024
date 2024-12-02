package main

import (
	"testing"
)

var testInput string = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestPuzzleOne(t *testing.T) {
	result, _ := puzzleOne(testInput)
	if result != 2 {
		t.Errorf("Result was not correct. Got: %d Want: %d", result, 2)
	}
}
