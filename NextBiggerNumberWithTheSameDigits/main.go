package main

import (
	"fmt"
	"strconv"
	"strings"
)

func NextBigger(n int) int {
	new := strconv.Itoa(n)
	if len(new) == 1 {
		return -1
	}
	total := ""
	result := 0
	lst := []string{}
	for _, v := range new {
		lst = append(lst, string(v))
	}
	for i := len(lst) - 1; i > 0; i-- {
		if lst[i] > lst[i-1] {
			lst[i], lst[i-1] = lst[i-1], lst[i]
			total = strings.Join(lst, "")
			result, _ = strconv.Atoi(total)
			return result
		}
	}
	return -1
}

func main() {
	fmt.Println(NextBigger(22))
}
