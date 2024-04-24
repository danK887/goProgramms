package main

import (
	"fmt"
)

func main() {
	arr := [3]float64{71.8, 56.2, 89.5}
	total := 0.0
	result := 0.0
	for _, v := range arr {
		total = total + v
		result = total / float64(len(arr))
	}
	fmt.Println(result)
}
