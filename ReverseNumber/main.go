package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(FindReverseNumber(18446744073709551615))
}

func FindReverseNumber(n uint64) uint64 {
	if n <= 1 {
		return 0
	}

	var count uint64 = 1

	// Для каждой длины числа генерируем палиндромы
	for digits := 1; ; digits++ {
		var palindromesInCurrentLength uint64

		if digits == 1 {
			palindromesInCurrentLength = 9
		} else if digits%2 == 0 {
			palindromesInCurrentLength = uint64(math.Pow10(digits/2-1) * 9)
		} else {
			palindromesInCurrentLength = uint64(math.Pow10(digits/2) * 9)
		}

		if count+palindromesInCurrentLength >= n {
			return findSpecificPalindrome(digits, n-count)
		}

		count += palindromesInCurrentLength
	}
}

func findSpecificPalindrome(digits int, position uint64) uint64 {
	var start uint64

	if digits == 1 {
		return position
	} else if digits%2 == 0 {
		start = uint64(math.Pow10(digits/2 - 1))
	} else {
		start = uint64(math.Pow10(digits / 2))
	}

	base := start + position - 1

	baseStr := strconv.FormatUint(base, 10)
	var reverseStr string

	if digits%2 == 0 {
		reverseStr = reverse(baseStr)
	} else {
		reverseStr = reverse(baseStr[:len(baseStr)-1])
	}

	palindromeStr := baseStr + reverseStr

	result, _ := strconv.ParseUint(palindromeStr, 10, 64)
	return result
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
