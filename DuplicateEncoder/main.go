package main

import (
	"fmt"
	"strings"
)

func DuplicateEncode(word string) string {
	result := ""
	countMap := make(map[rune]int)
	word = strings.ToLower(word)
	for _, let := range word {
		countMap[let]++
	}
	for _, let := range word {
		if countMap[let] != 1 {
			result = result + ")"
		} else {
			result = result + "("
		}
	}

	fmt.Println(countMap)
	return result
}

func main() {
	fmt.Println(DuplicateEncode("Success")) // )())())
}
