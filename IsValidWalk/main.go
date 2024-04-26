package main

import "fmt"

func IsValidWalk(walk []rune) bool {
	if len(walk) == 10 {
		sum := 0
		for i, _ := range walk {
			if walk[i] == 'n' {
				walk[i] = -'s'
			} else if walk[i] == 'w' {
				walk[i] = -'e'
			}
			sum += int(walk[i])
		}
		if sum == 0 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func main() {
	fmt.Println(IsValidWalk([]rune{'w'}))
}
