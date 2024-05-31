package main

import (
	"fmt"
	"strconv"
)

func Calc(s string) int {
	total1, val := "", ""
	lst1 := []string{}
	lst2 := []string{}
	sum1, sum2 := 0, 0
	gg := 0
	for _, v := range s {
		gg = int(v)
		val = strconv.Itoa(gg)
		total1 = total1 + val
	}

	for _, val := range total1 {
		lst1 = append(lst1, string(val))
	}

	for _, val := range lst1 {
		if val != "7" {
			lst2 = append(lst2, val)
		} else {
			lst2 = append(lst2, "1")
		}
	}

	for _, val := range lst1 {
		gg, _ = strconv.Atoi(val)
		sum1 = sum1 + gg
	}

	for _, val := range lst2 {
		gg, _ = strconv.Atoi(val)
		sum2 = sum2 + int(gg)
	}
	return sum1 - sum2
}

func main() {
	fmt.Println(Calc("aaaaaddddr")) //6
}
