package main

import "fmt"

func IsValidWalk(walk []rune) bool {
	if len(walk) != 10 {
		return false
	}
	for i, r := range walk {
		switch r {
		case 'n':
			walk[i] = -'s'
		case 'w':
			walk[i] = -'e'
		}
	}
	sum := 0
	for _, r := range walk {
		sum += int(r)
	}
	return sum == 0
}
func main() {
	fmt.Println(IsValidWalk([]rune{'n', 'e', 's', 'w', 'n', 's', 'w', 'n', 's'}))
}
