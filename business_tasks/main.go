package main

import (
	"fmt"
	"slices"
	"time"
)

var tests = [][]any{
	{[]string{"a", "b", "c", "d"}, 2},
	{[]string{"a", "b", "c", "d", "e"}, 3},
	{[]string{"alpha", "beta", "gamma", "delta", "epsilon"}, 1},
	{[]string{"a", "b"}, 1000},
	{[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}, 17},
	{[]string{"zlqamum", "yjsrpybmq", "tjllfea", "fxjqzznvg", "nvhekxr", "am", "skmazcey", "piklp", "olcqvhg", "dnpo", "bhcfc", "y", "h", "fj", "bjeoaxglt", "oafduixsz", "kmtbaxu", "qgcxjbfx", "my", "mlhy", "bt", "bo", "q"}, 9000000},
}

func main() {
	for i, test := range tests {
		start := time.Now()
		tasks := test[0].([]string)
		n := test[1].(int)
		fmt.Println(fmt.Sprintf("Test %d, tasks: %v, n: %d", i, tasks, n))
		fmt.Println(fmt.Sprintf("result: %s, execution time: %v", getTask(tasks, n, 0), time.Since(start)))
		fmt.Println("-----------------------------------------------------------")
	}
}

// recursive implementation
func getTask(tasks []string, n, startIndex int) string {
	if len(tasks) == 0 {
		return ""
	}

	if len(tasks) == 1 {
		return tasks[0]
	}

	// calculate index of the element to delete
	indexToDelete := startIndex + n - 1
	if len(tasks) < n+startIndex {
		// out of bounds index
		indexToDelete = indexToDelete - len(tasks)*(indexToDelete/len(tasks))
	}

	// set next start index
	if indexToDelete == len(tasks)-1 {
		// deleting last element, reset startIndex to 0
		startIndex = 0
	} else {
		startIndex = indexToDelete
	}

	// delete element
	newSlice := slices.Delete(tasks, indexToDelete, indexToDelete+1)

	return getTask(newSlice, n, startIndex)
}
