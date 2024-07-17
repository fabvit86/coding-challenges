package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"
)

var tests = [][]string{
	{"9 2 3", "4 8 7"},
	{"1 2", "4 5", "3 6"},
	{"1 1", "1 1"},
}

func main() {
	for i, test := range tests {
		start := time.Now()
		fmt.Println(fmt.Sprintf("Test %d, people: %v", i, test))
		fmt.Println(fmt.Sprintf("result: %v, execution time: %v", getPeople(test), time.Since(start)))
		fmt.Println("-----------------------------------------------------------")
	}
}

// return the "tallest-of-the-shortest" in each row and the "shortest-of-the-tallest" in each column
func getPeople(people []string) []int {
	var tallestOfTheShortest int
	tallestInColumns := make([]int, len(strings.Split(people[0], " ")))

	for _, row := range people {
		rowHeigths := strings.Split(row, " ")
		shortestInRow := math.MaxInt

		for j, rowHeigth := range rowHeigths {
			height, err := strconv.Atoi(rowHeigth)
			if err != nil {
				panic(err)
			}

			// update the tallest person in the column
			if height > tallestInColumns[j] {
				tallestInColumns[j] = height
			}

			// update the shortest person in the row
			if height < shortestInRow {
				shortestInRow = height
			}
		}

		// update the tallest of the shortest
		if shortestInRow > tallestOfTheShortest {
			tallestOfTheShortest = shortestInRow
		}
	}

	return []int{tallestOfTheShortest, slices.Min(tallestInColumns)}
}
