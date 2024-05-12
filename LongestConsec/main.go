package main

import (
	"fmt"
	"strings"
)

func LongestConsec(strarr []string, k int) string {
	if k <= 0 || len(strarr) == 0 || len(strarr) < k {
		return ""
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
	result = result[:len(result)-(k-1)]
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
	fmt.Println(LongestConsec([]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, 1)) // "abigailtheta"
}

// func concatenateEveryThreeStrings(arr []string) []string {
// 	var result []string
// 	z := 3
// 	for i := 0; i < len(arr); i++ {
// 		end := i + z
// 		if end > len(arr) {
// 			end = len(arr)
// 		}
// 		result = append(result, strings.Join(arr[i:end], ""))
// 	}

// 	return result
// }

// func main() {
// 	arr := []string{"zone", "abigail", "theta", "form", "libe", "zas"} // "abigailtheta"
// 	z := 3
// 	result := concatenateEveryThreeStrings(arr)

// 	fmt.Println(result[:len(result)-(z-1)])
// }
