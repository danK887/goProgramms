package main

import (
	"fmt"
)

func CountOnes(left, right uint64) uint64 {
	return countOnes(right) - countOnes(left-1)
}

func countOnes(n uint64) uint64 {
	if n == 0 {
		return 0
	}

	count := uint64(0)
	for i := uint64(1); i <= n; i <<= 1 {
		// Calculate the total pairs of complete cycles of 0s and 1s
		totalPairs := (n + 1) / (i << 1)
		count += totalPairs * i

		// Add the remaining 1s in the last incomplete cycle
		if remainder := (n + 1) % (i << 1); remainder > i {
			count += remainder - i
		}
	}
	return count
}

func main() {
	fmt.Println(CountOnes(12, 200000000000000)) // Example usage
}
