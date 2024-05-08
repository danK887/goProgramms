package main

import "fmt"

// m=109, z=122, a=97
func PrinterError(s string) string {
	countWrong := 0
	for _, l := range s {
		if byte(l) < 97 || byte(l) > 109 {
			countWrong++
		}
	}
	return fmt.Sprintf("%d/%d", countWrong, len(s))
}

func main() {
	fmt.Println(PrinterError("kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyzuuuuu"))
}
