package main

import "fmt"

func Calc(s string) int {
	sum := 0
	for _, v := range s {
		sum += int(v)
	}
	fmt.Println(sum)
	return 0
}

func main() {
	fmt.Println(Calc("abcdef")) //6
}
