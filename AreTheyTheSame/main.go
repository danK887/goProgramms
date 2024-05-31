package main

import (
	"fmt"
	"sort"
)

func Comp(array1 []int, array2 []int) bool {
	if array1 == nil || array2 == nil {
		return false
	}

	if len(array1) != len(array2) {
		return false
	}

	sort.Ints(array1)
	sort.Ints(array2)

	for i := range array1 {
		if (array1[i] * array1[i]) != array2[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Comp([]int{121, 144, 19, 161, 19, 144, 19, 11}, []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}))
}
