package main

import "testing"

var testInput string = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestPuzzleOne(t *testing.T) {
	result, _ := puzzleOne(testInput)
	if result != 18 {
		t.Errorf("Result was not correct. Got: %d Want: %d", result, 18)
	}
}

func TestPuzzleTwo(t *testing.T) {
	result, _ := puzzleTwo(testInput)
	if result != 9 {
		t.Errorf("Result was not correct. Got: %d Want: %d", result, 9)
	}
}
