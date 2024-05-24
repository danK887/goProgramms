package main

import (
	"fmt"
	"strings"
)

func ToNato(words string) string {
	alphabet := map[string]string{"A": "Alfa", "B": "Bravo", "C": "Charlie", "D": "Delta",
		"E": "Echo", "F": "Foxtrot", "G": "Golf", "H": "Hotel", "I": "India",
		"J": "Juliett", "K": "Kilo", "L": "Lima", "M": "Mike", "N": "November", "O": "Oscar",
		"P": "Papa", "Q": "Quebec", "R": "Romeo", "S": "Sierra", "T": "Tango",
		"U": "Uniform", "V": "Victor", "W": "Whiskey", "X": "Xray", "Y": "Yankee", "Z": "Zulu"}
	trimmedStr := strings.ReplaceAll(words, " ", "")
	strToConvert := strings.ToUpper(trimmedStr)
	wordArray := []string{}
	for _, let := range strToConvert {
		if _, ok := alphabet[string(let)]; ok {
			wordArray = append(wordArray, alphabet[string(let)])
		} else {
			wordArray = append(wordArray, string(let))
		}
	}
	result := strings.Join(wordArray, " ")
	return result
}

func main() {
	fmt.Println(ToNato("hubimjoc")) //India Foxtrot Yankee Oscar Uniform Charlie Alfa November Romeo Echo Alfa Delta"
}
