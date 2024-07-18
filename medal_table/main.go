package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"time"
)

var tests = [][]string{
	{"ITA JPN AUS", "KOR TPE UKR", "KOR KOR GBR", "KOR CHN TPE"},
	{"USA AUT ROM"},
	{"GER AUT SUI", "AUT SUI GER", "SUI GER AUT"},
}

func main() {
	for i, test := range tests {
		start := time.Now()
		fmt.Println(fmt.Sprintf("Test %d, input: %v", i, test))
		fmt.Println(fmt.Sprintf("result: %v, execution time: %v", generate(test), time.Since(start)))
		fmt.Println("-----------------------------------------------------------")
	}
}

func generate(results []string) []string {
	var medalTable []string
	medalTableMap := make(map[string][]int)

	// collect medals by country in a map
	for _, result := range results {
		podiumCountries := strings.Split(result, " ")
		if len(podiumCountries) != 3 {
			log.Fatal("invalid input")
		}

		for i, country := range podiumCountries {
			if entry, exists := medalTableMap[country]; !exists {
				medals := make([]int, 3)
				medals[i]++
				medalTableMap[country] = medals
			} else {
				entry[i]++
			}
		}
	}

	// sort and add to a slice
	for country, medals := range medalTableMap {
		newEntry := country + " " + strconv.Itoa(medals[0]) + " " + strconv.Itoa(medals[1]) + " " + strconv.Itoa(medals[2])

		if len(medalTable) == 0 {
			medalTable = append(medalTable, newEntry)
			continue
		}

		var added bool
		for i, entry := range medalTable {
			values := strings.Split(entry, " ")
			goldMedals, err := strconv.Atoi(values[1])
			if err != nil {
				log.Fatal(err)
			}
			silverMedals, err := strconv.Atoi(values[2])
			if err != nil {
				log.Fatal(err)
			}
			bronzeMedals, err := strconv.Atoi(values[3])
			if err != nil {
				log.Fatal(err)
			}

			if medals[0] > goldMedals ||
				(medals[0] == goldMedals && medals[1] > silverMedals) ||
				(medals[0] == goldMedals && medals[1] == silverMedals && medals[2] > bronzeMedals) ||
				(medals[0] == goldMedals && medals[1] == silverMedals && medals[2] == bronzeMedals && country < values[0]) {
				// insert before current element
				medalTable = slices.Insert(medalTable, i, newEntry)
				added = true
				break
			}
		}

		if !added {
			medalTable = append(medalTable, newEntry)
		}
	}

	return medalTable
}
