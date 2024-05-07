package main

import (
	"fmt"
	"strings"
)

func solve(str string) string {
	upInd := 0
	lowInd := 0
	// 65-90 Upper
	// 97-122 Lower
	for _, v := range str {
		if byte(v) >= 65 && byte(v) <= 90 {
			upInd += 1
		} else if byte(v) >= 97 && byte(v) <= 122 {
			lowInd += 1
		}
		if upInd > lowInd {
			str = strings.ToUpper(str)
		} else {
			str = strings.ToLower(str)
		}
	}

	return str
}

func main() {
	fmt.Println(solve("HELo"))
}
