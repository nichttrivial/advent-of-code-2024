package main

import (
	"log"
	"regexp"
	"slices"
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
	pd := newPuzzleDataFromString(input)
	return pd.countXmas(), nil
}

func puzzleTwo(input string) (int, error) {
	pd := newPuzzleDataFromString(input)
	return pd.countXedmas(), nil
}

type puzzleData struct {
	rows  [][]string
	cols  [][]string
	diags [][]string
}

func newPuzzleDataFromString(input string) puzzleData {
	pd := puzzleData{}

	//Build rows from input string
	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		row := strings.Split(line, "")
		pd.rows = append(pd.rows, row)
	}

	//Build coloumns from rows. Assume all rows have the same length
	for i := 0; i < len(pd.rows[0]); i++ {
		col := make([]string, len(pd.rows))
		for j := 0; j < len(pd.rows); j++ {
			n := pd.rows[j][i]
			col[j] = n
		}
		slices.Reverse(col) //Had to draw this to get this is important to build the diagonals from the columns...
		pd.cols = append(pd.cols, col)
	}

	//Build diagonals from rows
	for i := 0; i < len(pd.rows); i++ {
		diag := make([]string, len(pd.rows)-i)
		for j := 0; j < len(pd.rows)-i; j++ {
			n := pd.rows[i+j][j]
			diag[j] = n
		}
		pd.diags = append(pd.diags, diag)
	}
	for i := 1; i < len(pd.rows); i++ {
		diag := make([]string, len(pd.rows)-i)
		for j := 0; j < len(pd.rows)-i; j++ {
			n := pd.rows[j][i+j]
			diag[j] = n
		}
		pd.diags = append(pd.diags, diag)
	}

	//Build diagonals from cols
	for i := 0; i < len(pd.cols); i++ {
		diag := make([]string, len(pd.cols)-i)
		for j := 0; j < len(pd.cols)-i; j++ {
			n := pd.cols[i+j][j]
			diag[j] = n
		}
		pd.diags = append(pd.diags, diag)
	}
	for i := 1; i < len(pd.cols); i++ {
		diag := make([]string, len(pd.cols)-i)
		for j := 0; j < len(pd.cols)-i; j++ {
			n := pd.cols[j][i+j]
			diag[j] = n
		}
		pd.diags = append(pd.diags, diag)
	}

	return pd
}

func (pd puzzleData) countXmas() int {
	re := regexp.MustCompile("XMAS")
	result := 0
	for _, row := range pd.rows {
		result += len(re.FindAllString(strings.Join(row, ""), -1))
		slices.Reverse(row)
		result += len(re.FindAllString(strings.Join(row, ""), -1))
	}
	for _, col := range pd.cols {
		result += len(re.FindAllString(strings.Join(col, ""), -1))
		slices.Reverse(col)
		result += len(re.FindAllString(strings.Join(col, ""), -1))
	}
	for _, diag := range pd.diags {
		result += len(re.FindAllString(strings.Join(diag, ""), -1))
		slices.Reverse(diag)
		result += len(re.FindAllString(strings.Join(diag, ""), -1))
	}
	return result
}

func (pd puzzleData) countXedmas() int {
	result := 0
	for i, row := range pd.rows {
		for j, char := range row {
			if char == "A" {
				if i-1 < 0 || j-1 < 0 || i+1 >= len(pd.rows) || j+1 >= len(pd.rows) {
					continue
				}
				if (pd.rows[i-1][j+1] == "S" && pd.rows[i+1][j-1] == "M" ||
					pd.rows[i-1][j+1] == "M" && pd.rows[i+1][j-1] == "S") &&
					(pd.rows[i-1][j-1] == "S" && pd.rows[i+1][j+1] == "M" ||
						pd.rows[i-1][j-1] == "M" && pd.rows[i+1][j+1] == "S") {
					result += 1
				}
			}
		}
	}
	return result
}
