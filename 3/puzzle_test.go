package main

import (
	"testing"
)

func TestPuzzleOne(t *testing.T) {
	testInput := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	result, _ := puzzleOne(testInput)
	if result != 161 {
		t.Errorf("Result was not correct. Got: %d Want: %d", result, 151)
	}
}

func TestPuzzleTwo(t *testing.T) {
	testInput := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	result, _ := puzzleTwo(testInput)
	if result != 48 {
		t.Errorf("Result was not correct. Got: %d Want: %d", result, 48)
	}
}
