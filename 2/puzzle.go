package main

import (
	"cmp"
	"log"
	"regexp"
	"slices"
	"strconv"
	"strings"

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
	re := regexp.MustCompile("[0-9]+")
	cnt := 0

out:
	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		reportString := re.FindAllString(line, -1)

		reportInt := make([]int, len(reportString))
		for i, num := range reportString {
			n, _ := strconv.Atoi(num)
			reportInt[i] = n
		}

		asc := func(a, b int) int { return cmp.Compare(a, b) }
		desc := func(a, b int) int { return -1 * cmp.Compare(a, b) }

		if !slices.IsSortedFunc(reportInt, asc) && !slices.IsSortedFunc(reportInt, desc) {
			continue
		}

		for i := 0; i < len(reportInt)-1; i++ {
			if reportInt[i]-reportInt[i+1] > 3 ||
				reportInt[i]-reportInt[i+1] < -3 ||
				reportInt[i]-reportInt[i+1] == 0 {
				continue out
			}
		}

		cnt += 1
	}

	return cnt, nil
}

func puzzleTwo(input string) (int, error) {
	panic("unimplemented")
}
