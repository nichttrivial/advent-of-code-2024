package utility

import (
	"errors"
	"os"
)

func ReadInputToString() (string, error) {
	if len(os.Args) < 2 {
		return "", errors.New("no path submitted")
	}

	inputPath := os.Args[1]
	input, err := os.ReadFile(inputPath)
	if err != nil {
		return "", errors.New("no file at " + inputPath)
	}

	return string(input), nil
}

func AbsInt(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
