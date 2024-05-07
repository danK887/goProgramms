package main

import "fmt"

func QuarterOf(month int) int {
	switch {
	case month <= 3:
		return 1
	case month <= 6:
		return 2
	case month <= 9:
		return 3
	case month <= 12:
		return 4
	}
	return 0
}

func main() {
	fmt.Println(QuarterOf(11))
}
