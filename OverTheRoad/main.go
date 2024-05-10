package main

import "fmt"

func OverTheRoad(address int, n int) int {
	i, j := n, 0
	street := make(map[int]int)
	for i > 0 && j < n {
		street[2*i] = j*2 + 1
		i--
		j++
	}
	fmt.Println(street)
	for key, val := range street {
		if val == address {
			return key
		}
	}
	return street[address]
}

func main() {
	fmt.Println(OverTheRoad(1, 3)) //6
}

// 1|   |8
// 3|   |6
// 5|   |4
// 7|   |2
//   you
