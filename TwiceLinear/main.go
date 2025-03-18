package main

import (
	"fmt"
)

func DblLinear(n int) int {
	sequence := []int{1}
	yList, zList := []int{}, []int{}

	for i := 0; i < n; i++ {
		x := sequence[i]
		y, z := 2*x+1, 3*x+1

		yList = append(yList, y)
		zList = append(zList, z)

		if len(yList) > 0 && (len(zList) == 0 || yList[0] < zList[0]) {
			sequence = append(sequence, yList[0])
			yList = yList[1:]
		} else {
			sequence = append(sequence, zList[0])
			zList = zList[1:]
		}
	}

	return sequence[n]
}

func main() {
	fmt.Println(DblLinear(100000))
}
