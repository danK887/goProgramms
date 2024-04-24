package main

import "fmt"

func IsTriangle(a, b, c int) bool {
	if a+b > c && a > 0 && b > 0 && c > 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(IsTriangle(0, 6, 3))
}
