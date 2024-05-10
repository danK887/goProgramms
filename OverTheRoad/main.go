package main

import "fmt"

// my code @no optimize :(
// func OverTheRoad(address int, n int) int {
// 	i, j := n, 0
// 	street := make(map[int]int)
// 	for i > 0 && j < n {
// 		street[2*i] = j*2 + 1
// 		i--
// 		j++
// 	}
// 	fmt.Println(street)
// 	for key, val := range street {
// 		if val == address {
// 			return key
// 		}
// 	}
// 	return street[address]
// }

// computer code
func OverTheRoad(address int, n int) int {
	return 2*n - address + 1
}

func main() {
	fmt.Println(OverTheRoad(6, 4)) //6
}

// 1|   |8
// 3|   |6
// 5|   |4
// 7|   |2
//   you
