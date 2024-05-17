package main

import (
	"fmt"
	"strings"
)

func Accum(s string) string {
	result := []string{}
	letStr := ""
	total := []string{}
	capital := ""
	counter := 0
	for _, c := range s {
		letStr = ""
		for i := 0; i <= counter; i++ {
			letStr += string(c)
		}
		total = append(total, letStr)
		counter++
	}
	for _, word := range total {
		capital = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		result = append(result, capital)
	}
	//return capital
	return strings.Join(result, "-")
}

func main() {
	fmt.Println(Accum("ZpglnRxqenU")) //Z-Pp-Ggg-Llll-Nnnnn-Rrrrrr-Xxxxxxx-Qqqqqqqq-Eeeeeeeee-Nnnnnnnnnn-Uuuuuuuuuuu
}
