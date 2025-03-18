package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{3, 2, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if idx, found := numMap[complement]; found {
			return []int{idx, i}
		}
		numMap[num] = i
	}
	return []int{}
}
