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
	pd := newPuzzleDataFromString(input)
	return pd.countMPN(), nil
}

func puzzleTwo(input string) (int, error) {
	return 0, nil
}

type puzzleData struct {
	orders  map[int][]int
	updates [][]int
}

func newPuzzleDataFromString(input string) puzzleData {
	pd := puzzleData{
		orders: make(map[int][]int),
	}

	orderRe := regexp.MustCompile(`[0-9]+\|[0-9]+`)
	orders := orderRe.FindAllString(input, -1)
	for _, order := range orders {
		nums := strings.Split(order, "|")
		l, _ := strconv.Atoi(nums[0])
		r, _ := strconv.Atoi(nums[1])
		pd.orders[l] = append(pd.orders[l], r)
	}

	updateRe := regexp.MustCompile(`([0-9]+,)+[0-9]+`)
	updates := updateRe.FindAllString(input, -1)
	for _, update := range updates {
		nums := strings.Split(update, ",")
		numsInt := make([]int, len(nums))
		for i, num := range nums {
			numInt, _ := strconv.Atoi(num)
			numsInt[i] = numInt
		}
		pd.updates = append(pd.updates, numsInt)
	}

	return pd
}

func (pd puzzleData) validUpdateIdxs() []int {
	var idxs []int
start:
	for i, update := range pd.updates {
		for j, num := range update {
			n := j + 1
			for ; n < len(update); n++ {
				if slices.Contains(pd.orders[update[n]], num) {
					continue start
				}
			}
		}
		idxs = append(idxs, i)
	}
	return idxs
}

func (pd puzzleData) countMPN() int {
	idxs := pd.validUpdateIdxs()
	cnt := 0

	for _, idx := range idxs {
		middleIdx := len(pd.updates[idx]) / 2
		cnt += pd.updates[idx][middleIdx]
	}

	return cnt
}
