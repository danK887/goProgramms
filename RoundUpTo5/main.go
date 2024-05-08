package main

import "fmt"

func RoundToNext5(n int) int {
	for n%5 != 0 {
		n++
	}
	return n
}

func main() {
	fmt.Println(RoundToNext5(-5))
}
