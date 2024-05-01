package main

import "fmt"

func FakeBin(x string) string {
	result := ""
	for i, _ := range x {
		if x[i] == '5' || x[i] == '6' || x[i] == '7' || x[i] == '8' || x[i] == '9' {
			result += "1"
		} else {
			result += "0"
		}
	}
	return result
}

func main() {
	fmt.Println(FakeBin("800857237867"))
}
