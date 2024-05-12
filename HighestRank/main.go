package main

import "fmt"

func HighestRank(nums []int) int {
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}
	//fmt.Println(countMap)
	countValue := 0
	hightNum := 0
	for key, val := range countMap {
		if val > countValue {
			countValue = val
			hightNum = key
		} else if val == countValue {
			if key > hightNum {
				hightNum = key
			}
		}
	}
	return hightNum
}

func main() {
	fmt.Println(HighestRank([]int{2, 1, 5, 3, 9, 7, 8, 6, 4})) //12
}
