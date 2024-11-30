package main

import (
	"fmt"

	input "github.com/nichttrivial/advent-of-code-2024/internal"
)

func main() {
	input, err := input.ReadInputToString()
	if err != nil {
		fmt.Println(err.Error())
	}

	puzzle(input)
}

func puzzle(input string) int {
	fmt.Println(string(input))
	return 0
}
