package main

import (
	"fmt"
)

func beeramid(bonus int, price float64) int {
	totalCans := int(float64(bonus) / price)
	level := 0
	cansNeeded := 0

	for {
		level++
		cansNeeded = level * level // Calculate cans needed for the current level
		if totalCans >= cansNeeded {
			totalCans -= cansNeeded // Deduct the cans used for this level
		} else {
			break // Can't build the next level
		}
	}

	return level - 1 // Return the number of complete levels
}

func main() {
	fmt.Println(beeramid(1500, 2.0)) // should return 12
	fmt.Println(beeramid(5000, 3.0)) // should return 16
}
