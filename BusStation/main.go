package main

import "fmt"

func Number(busStops [][2]int) int {
	totalPeople := 0
	for _, v1 := range busStops {
		totalPeople = totalPeople + v1[0] - v1[1]
	}

	return totalPeople // Good Luck!
}

func main() {
	fmt.Println(Number([][2]int{{0, 0}}))
}
