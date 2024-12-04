package main

import "testing"

var testInput string = `....XXMAS.
.SAMXMS...
...S..A...
..A.A.MS.X
XMASAMX.MM
X.....XA.A
S.S.S.S.SS
.A.A.A.A.A
..M.M.M.MM
.X.X.XMASX`

func TestPuzzleOne(t *testing.T) {
	result, _ := puzzleOne(testInput)
	if result != 18 {
		t.Errorf("Result was not correct. Got: %d Want: %d", result, 18)
	}
}
