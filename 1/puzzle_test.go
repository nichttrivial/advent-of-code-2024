package main

import (
	"reflect"
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

func TestNewPuzzleDataFromStringSuccess(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  puzzleData
	}{
		{"Success 1", "1   2", puzzleData{right: []int{2}, left: []int{1}}},
		{"Success 2", "1   2\n3   4", puzzleData{right: []int{2, 4}, left: []int{1, 3}}},
		{"Success 3", "1   2\n3   4\n", puzzleData{right: []int{2, 4}, left: []int{1, 3}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out, _ := newPuzzleDataFromString(test.input)
			if !reflect.DeepEqual(out, test.want) {
				t.Errorf("structs not equal")
			}
		})
	}
}

func TestNewPuzzleDataFromStringErrors(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  puzzleData
	}{
		{"Error 1", "1   n", puzzleData{}},
		{"Error 2", "1   2\nasd", puzzleData{}},
		{"Error 3", "asdfasdf", puzzleData{}},
		{"Error 4", "", puzzleData{}},
		{"Error 5", "1   2   4", puzzleData{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out, err := newPuzzleDataFromString(test.input)
			if !reflect.DeepEqual(out, test.want) {
				t.Errorf("structs not equal")
			}
			if err == nil {
				t.Errorf("No error on invalid input")
			}
		})
	}
}
