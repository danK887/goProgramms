package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func ValidISBN10(isbn string) bool {
	sum := 0
	num := 0
	if len(isbn) != 10 {
		return false
	}
	for _, val := range isbn[:9] {
		if !unicode.IsDigit(rune(val)) {
			return false
		}
	}
	for ind, val := range isbn {
		if val == 'X' || val == 'x' {
			num = 10
		} else {
			num, _ = strconv.Atoi(string(val))
		}
		sum = sum + ((ind + 1) * num)
	}
	return sum%11 == 0
}

func main() {
	fmt.Println(ValidISBN10("048665088x"))

}
