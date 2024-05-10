package main

import (
	"fmt"
	"slices"
	"strings"
	"time"
)

var (
	tests = [][]any{
		{"aaaa", "zzzz", []string{"a a a z", "a a z a", "a z a a", "z a a a", "a z z z", "z a z z", "z z a z", "z z z a"}},
		{"aaaa", "bbbb", []string{}},
		{"aaaa", "mmnn", []string{}},
		{"aaaa", "zzzz", []string{"bz a a a", "a bz a a", "a a bz a", "a a a bz"}},
		{"aaaa", "zzzz", []string{"cdefghijklmnopqrstuvwxyz a a a", "a cdefghijklmnopqrstuvwxyz a a", "a a cdefghijklmnopqrstuvwxyz a", "a a a cdefghijklmnopqrstuvwxyz"}},
		{"aaaa", "bbbb", []string{"b b b b"}},
		{"zzzz", "aaaa", []string{"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk",
			"abcdefghijkl abcdefghijkl abcdefghijkl abcdefghijk"}},
	}
)

type wordWithCost struct {
	word   string
	cost   int
	parent string
}

func main() {
	for i, test := range tests {
		input := test[0].(string)
		output := test[1].(string)
		forbiddenWords := test[2].([]string)
		start := time.Now()

		fmt.Println("Test", i, "input:", input, "output:", output)

		steps := findPath(input, output, forbiddenWords)
		if steps < 0 {
			fmt.Println("Test", i, "no possible solutions found. Input:", input, "Output:", output, "execution time:", time.Since(start))
		} else {
			fmt.Println("Test", i, "shortest path:", steps, "steps. Input:", input, "Output:", output, "execution time:", time.Since(start))
		}

		fmt.Println("-----------------------------------------------------------")
	}
}

func findPath(input, output string, forbiddenWords []string) int {
	if input == output {
		return 0
	}

	var iterations int
	examined := make(map[string]string)
	queue := []wordWithCost{{
		word:   input,
		cost:   0,
		parent: "",
	}}

	for {
		iterations++

		if len(queue) == 0 {
			fmt.Println("iterations:", iterations)
			return -1
		}

		// pop element from queue and add to examined set
		currentWord := queue[0]
		queue = queue[1:]
		examined[currentWord.word] = currentWord.parent

		// check if target is reached
		if currentWord.word == output {
			// retrieve path
			path := make([]string, 0)
			path = append(path, currentWord.word)
			parent := examined[currentWord.word]
			for parent != "" {
				path = append([]string{parent}, path...)
				parent = examined[parent]
			}

			fmt.Println("iterations:", iterations)
			fmt.Println("shortest path:", strings.Join(path, " -> "))

			return currentWord.cost
		}

		// collect all possible neighbors of currentWord
		neighbors := getNeighbors(currentWord.word, forbiddenWords)

		// check all neighbors
		for _, neighbor := range neighbors {
			if _, contains := examined[neighbor]; !contains {
				var inQueue bool
				for _, element := range queue {
					if element.word == neighbor {
						inQueue = true
						break
					}
				}
				if !inQueue {
					queue = append(queue, wordWithCost{
						word:   neighbor,
						cost:   currentWord.cost + 1,
						parent: currentWord.word,
					})
				}
			}
		}

	}
}

// a word is forbidden if it's a combination of 4 letters contained in the forbid slice, one letter per index
func wordIsForbidden(word []rune, forbiddenWords []string) bool {
	for _, s := range forbiddenWords {
		forbidden := true
		forbiddenLetters := strings.Split(s, " ")
		for i, letter := range word {
			if !slices.Contains([]rune(forbiddenLetters[i]), letter) {
				forbidden = false
				break
			}
		}

		if forbidden {
			return true
		}
	}

	return false
}

func nextLetter(letter rune) rune {
	if letter == 'z' {
		return 'a'
	}

	return letter + 1
}

func previousLetter(letter rune) rune {
	if letter == 'a' {
		return 'z'
	}

	return letter - 1
}

// return all possible neighbors of the given word
func getNeighbors(word string, forbiddenWords []string) []string {
	var neighbors []string
	wordRunes := []rune(word)
	for i, letter := range wordRunes {
		neighborFw := append([]rune(nil), wordRunes...)
		neighborFw[i] = nextLetter(letter)
		if !wordIsForbidden(neighborFw, forbiddenWords) {
			neighbors = append(neighbors, string(neighborFw))
		}

		neighborBw := append([]rune(nil), wordRunes...)
		neighborBw[i] = previousLetter(letter)
		if !wordIsForbidden(neighborBw, forbiddenWords) {
			neighbors = append(neighbors, string(neighborBw))
		}
	}

	return neighbors
}
