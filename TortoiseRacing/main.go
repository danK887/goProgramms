package main

import "fmt"

func Race(v1, v2, g int) [3]int {
	if v1 >= v2 {
		return [3]int{-1, -1, -1}
	}
	totalTime := float64(g) / float64(v2-v1)
	hours := int(totalTime)
	minutes := int((totalTime - float64(hours)) * 60)
	seconds := int(((totalTime-float64(hours))*60 - float64(minutes)) * 60)

	return [3]int{hours, minutes, seconds}
}

func main() {
	fmt.Println("Race1: ", Race(80, 91, 37))
	fmt.Println("Race2: ", Race2(80, 91, 37))
}

func Race2(v1, v2, g int) [3]int {
	if v1 >= v2 {
		return [3]int{-1, -1, -1}
	}

	t := int(float64(g) / float64(v2-v1) * 3600.0)

	h := t / 3600
	t = t % 3600

	m := t / 60
	s := t % 60

	return [3]int{h, m, s}
}
