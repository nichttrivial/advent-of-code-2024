package main

import (
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
	result, err = puzzleTwo(input)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(result)
}

func puzzleOne(input string) (int, error) {
	return analyzeReports(input, false)
}
func puzzleTwo(input string) (int, error) {
	return analyzeReports(input, true)
}

func analyzeReports(input string, dampener bool) (int, error) {
	re := regexp.MustCompile("[0-9]+")
	cnt := 0

	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {

		reportInt := reportToIntSlice(line, re)
		if isReportSafe(reportInt, dampener) {
			cnt += 1
		}
	}

	return cnt, nil
}

func isReportSafe(report []int, dampener bool) bool {
	var cmp func(a, b int) bool
	if report[0] == report[1] {
		if dampener {
			return isReportSafe(slices.Delete(append([]int(nil), report...), 0, 1), false)
		}
		return false
	} else if report[0] < report[1] {
		cmp = func(a, b int) bool { return a < b }
	} else {
		cmp = func(a, b int) bool { return a > b }
	}

	for i := 0; i < len(report)-1; i++ {
		if !cmp(report[i], report[i+1]) ||
			utility.AbsInt(report[i]-report[i+1]) > 3 {
			if dampener {
				for i := 0; i < len(report); i++ {
					if isReportSafe(slices.Delete(append([]int(nil), report...), i, i+1), false) {
						return true
					}
				}
			}
			return false
		}
	}

	return true
}

func reportToIntSlice(line string, re *regexp.Regexp) []int {
	reportString := re.FindAllString(line, -1)

	reportInt := make([]int, len(reportString))
	for i, num := range reportString {
		n, _ := strconv.Atoi(num)
		reportInt[i] = n
	}

	return reportInt
}
