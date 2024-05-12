package main

import (
	"fmt"
	"strings"
)

func LongestConsec(strarr []string, k int) string {
	if k <= 0 || len(strarr) == 0 || len(strarr) < k {
		return ""
	}
	if k == 1 {
		longestWord := ""
		maxLength := 0

		for _, word := range strarr {
			if len(word) > maxLength {
				maxLength = len(word)
				longestWord = word
			}
		}

		return longestWord
	}
	total := ""
	var result []string
	for i := 0; i < len(strarr); i++ {
		end := i + k
		if end > len(strarr) {
			end = len(strarr)
		}
		result = append(result, strings.Join(strarr[i:end], ""))
	}
	lenStr := 0
	for i := 0; i < len(result)-1; i++ {
		if len(result[i]) > lenStr {
			lenStr = len(result[i])
			total = result[i]
		}
	}
	return total
}

func main() {
	fmt.Println(LongestConsec([]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, 1)) 
}

