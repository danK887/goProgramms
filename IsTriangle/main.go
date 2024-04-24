package main

import (
	"fmt"
)

func IsTriangle(a, b, c int) bool {
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}
	return a+b > c && a > 0 && b > 0 && c > 0
}

func main() {
	fmt.Println(IsTriangle(2, 5, 1))
}
