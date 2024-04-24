package main

import "fmt"

func Disemvowel(comment string) string {
	result := ""
	vowelArray := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for _, vowel := range comment {
		for _, vowelArrayElement := range vowelArray {
			if vowel == vowelArrayElement {
				result += string(vowel)
			}
		}
	}
	return result
}

func main() {
	fmt.Println(Disemvowel("This website is for losers LOL!"))
}
