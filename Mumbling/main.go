package main

import (
	"fmt"
	"strings"
)

func Accum(s string) string {
	result := []string{}
	letStr := ""
	counter := 0
	for _, c := range s {
		letStr = ""
		for i := 0; i <= counter; i++ {
			letStr += string(c)
		}
		result = append(result, letStr)
		counter++
	}
	//fmt.Println(result)
	return strings.Join(result, "-")
}

func main() {
	fmt.Println(Accum("ZpglnRxqenU")) //Z-Pp-Ggg-Llll-Nnnnn-Rrrrrr-Xxxxxxx-Qqqqqqqq-Eeeeeeeee-Nnnnnnnnnn-Uuuuuuuuuuu
}
