package main

import "fmt"

func SumCubes(n int) int {
	sum := 0
	for n >= 1 {
		sum += n * n * n
		n -= 1
	}
	return sum
}

func main() {
	fmt.Println(SumCubes(3))
}
