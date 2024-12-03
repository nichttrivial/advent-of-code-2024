package main

import (
	"log"
	"regexp"
	"strconv"

	utility "github.com/nichttrivial/advent-of-code-2024/internal"
)

func main() {
	input, err := utility.ReadInputToString()
	if err != nil {
		log.Fatalln(err.Error())
	}

	//Part 1
	result, err := puzzleOne(input)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(result)

	//Part 2
	// result, err = puzzleTwo(input)
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// }
	// log.Println(result)
}

func puzzleOne(input string) (int, error) {
	re := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	instructions := re.FindAllString(input, -1)
	return sumSlice(execInstructions(instructions)), nil
}

func execInstructions(instructions []string) []int {
	re := regexp.MustCompile("[0-9]+")
	result := make([]int, len(instructions))

	for i, instruction := range instructions {
		factors := re.FindAllString(instruction, 2)
		a, _ := strconv.Atoi(factors[0])
		b, _ := strconv.Atoi(factors[1])
		result[i] = a * b
	}
	return result
}

func sumSlice(s []int) int {
	sum := 0
	for _, n := range s {
		sum += n
	}
	return sum
}
