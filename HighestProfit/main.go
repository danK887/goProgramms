package main

import "fmt"

func MinMax(arr []int) [2]int {
	min := arr[0]
	max := arr[0]

	for _, num := range arr {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return [2]int{min, max}
}

func main() {
	fmt.Println(MinMax([]int{665, -21, 1, 2, 3, 4, 5, 0, -3463}))
}
