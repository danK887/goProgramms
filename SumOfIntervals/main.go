package main

import (
	"fmt"
)

func main() {
	a := [][2]int{{1, 2}, {6, 10}, {11, 15}} //9
	//b := [][2]int{{1, 5}, {10, 20}, {15, 18}}
	fmt.Println(SumOfIntervals(a))
}

func SumOfIntervals(intervals [][2]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sortIntervals(intervals)

	start := intervals[0][0]
	end := intervals[0][1]
	sum := 0

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= end {
			if intervals[i][1] > end {
				end = intervals[i][1]
			}
		} else {
			sum += end - start
			start = intervals[i][0]
			end = intervals[i][1]
		}
	}

	sum += end - start

	return sum
}

func sortIntervals(intervals [][2]int) {
	for i := 0; i < len(intervals); i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[i][0] > intervals[j][0] {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}
}
