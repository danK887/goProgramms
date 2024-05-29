package main

import (
	"fmt"
	"strings"
)

func duplicate_count(s1 string) int {
	s1 = strings.ToLower(s1)
	countMap := make(map[rune]int)
	counter := 0
	for _, ch := range s1 {
		countMap[ch]++
	}
	for k := range countMap {
		if countMap[k] > 1 {
			counter++
		}
	}
	return counter
}

func main() {
	fmt.Println(duplicate_count("abcdeaB11"))
}
