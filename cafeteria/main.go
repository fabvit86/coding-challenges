package main

import (
	"fmt"
	"time"
)

// 12 hours time format (HH:MM)
const layout = "03:04"

var tests = [][][]int{
	{
		{9},
		{1},
		{1},
	},
	{
		{6},
		{9},
		{120},
	},
	{
		{6, 9},
		{9, 10},
		{120, 121},
	},
	{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		{190, 190, 190, 190, 190, 190, 190, 190, 190, 190},
	},
	{
		{7, 4, 0, 0, 2, 1, 6, 7, 7, 0, 8, 6, 0, 5, 0, 6, 7, 9, 0, 2, 4, 8, 4, 7, 9, 2, 4, 4, 3, 1, 4, 5, 8, 8, 2, 5, 7, 8, 7, 5, 6, 8, 8, 0, 1, 3, 5, 0, 8},
		{26, 14, 1, 4, 16, 28, 16, 6, 4, 5, 21, 18, 5, 2, 21, 21, 28, 22, 5, 22, 26, 16, 14, 19, 19, 19, 4, 12, 24, 4, 30, 16, 28, 20, 25, 2, 30, 18, 4, 6, 9, 22, 8, 3, 7, 29, 8, 30, 6},
		{151, 264, 280, 89, 63, 57, 15, 120, 28, 296, 76, 269, 90, 106, 31, 222, 291, 52, 102, 73, 140, 248, 44, 187, 76, 49, 296, 106, 54, 119, 54, 283, 263, 285, 275, 127, 108, 82, 84, 241, 169, 203, 244, 256, 109, 288, 9, 262, 103},
	},
}

func main() {
	for i, test := range tests {
		start := time.Now()
		fmt.Println(fmt.Sprintf("Test %d, input:%v\nlatestTime: %v\nexecution time: %v",
			i, test, latestTime(test[0], test[1], test[2]), time.Since(start)))
		fmt.Println("-----------------------------------------------------------")
	}
}

func latestTime(offset, walkingTime, drivingTime []int) string {
	bestLatestTime := time.Date(1999, time.January, 1, 0, 0, 0, 0, time.UTC)
	tArrive := time.Date(2000, time.January, 1, 14, 30, 0, 0, time.UTC)

	for i := range offset {
		tMax := tArrive.Add(-time.Minute * time.Duration(drivingTime[i]))
		tDep := findDepartureTime(tMax, offset[i])
		tLatest := tDep.Add(-time.Minute * time.Duration(walkingTime[i]))
		if tLatest.After(bestLatestTime) {
			bestLatestTime = tLatest
		}
	}

	return bestLatestTime.Format(layout)
}

// findDepartureTime returns the latest possible departure time
func findDepartureTime(tMax time.Time, offset int) time.Time {
	tDep := time.Date(2000, time.January, 1, tMax.Hour(), offset, 0, 0, time.UTC)
	if tDep.After(tMax) {
		tDep = tDep.Add(-time.Hour)
	}
	for {
		if tDep.Add(time.Minute * 10).After(tMax) {
			return tDep
		}
		tDep = tDep.Add(time.Minute * 10)
	}
}
