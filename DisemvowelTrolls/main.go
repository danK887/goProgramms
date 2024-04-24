package main

import (
	"fmt"
	"strings"
)

func Disemvowel(comment string) string {
	result := ""
	vowelString := "aeiouAEIOU"
	for _, char := range comment {
		if !strings.ContainsRune(vowelString, char) {
			result += string(char)
		}
	}
	return result
}

func main() {
	fmt.Println(Disemvowel("This website is for losers LOL!"))
}
