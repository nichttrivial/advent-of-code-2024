package main

import (
	"testing"
)

var testInput string = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

func TestPuzzleOne(t *testing.T) {
	result, _ := puzzleOne(testInput)
	if result != 161 {
		t.Errorf("Result was not correct. Got: %d Want: %d", result, 151)
	}
}
