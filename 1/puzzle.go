package main

import (
	"errors"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"

	input "github.com/nichttrivial/advent-of-code-2024/internal"
)

func main() {
	input, err := input.ReadInputToString()
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
	result, err = puzzleTwo(input)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(result)
}

func puzzleOne(input string) (int, error) {
	result := 0
	pd, err := newPuzzleDataFromString(input)
	if err != nil {
		return result, err
	}

	sort.Slice(pd.left, func(i, j int) bool {
		return pd.left[i] < pd.left[j]
	})

	sort.Slice(pd.right, func(i, j int) bool {
		return pd.right[i] < pd.right[j]
	})

	for i := 0; i < len(pd.right); i++ {
		diff := pd.left[i] - pd.right[i]
		if diff < 0 {
			diff = -1 * diff
		}
		result += diff
	}

	return result, nil
}

func puzzleTwo(input string) (int, error) {
	result := 0
	pd, err := newPuzzleDataFromString(input)
	if err != nil {
		return result, err
	}

	for _, l := range pd.left {
		cnt := 0
		for _, r := range pd.right {
			if l == r {
				cnt += 1
			}
		}
		result += l * cnt
	}

	return result, nil
}

type puzzleData struct {
	right []int
	left  []int
}

func newPuzzleDataFromString(input string) (puzzleData, error) {
	pd := puzzleData{}
	re := regexp.MustCompile("[0-9]+")

	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		numbers := re.FindAllString(line, 2)

		l, err := strconv.Atoi(numbers[0])
		if err != nil {
			return pd, errors.New("error in the left numbers")
		}

		r, err := strconv.Atoi(numbers[1])
		if err != nil {
			return pd, errors.New("error in the right numbers")
		}

		pd.left = append(pd.left, l)
		pd.right = append(pd.right, r)
	}

	if pd.left == nil || pd.right == nil || len(pd.right) != len(pd.left) {
		return pd, errors.New("invalid input data")
	}

	return pd, nil
}
