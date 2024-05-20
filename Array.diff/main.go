package main

import "fmt"

func ArrayDiff(a, b []int) []int {
	bMap := make(map[int]bool)
	for _, val := range b {
		bMap[val] = true
	}

	var result []int
	for _, val := range a {
		if !bMap[val] {
			result = append(result, val)
		}
	}

	return result
}

func main() {
	fmt.Println(ArrayDiff([]int{1, 2, 2, 1, 3, 4}, []int{1, 2}))
}
