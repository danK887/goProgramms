package main

import "fmt"

func Evaporator(content float64, evapPerDay int, threshold int) int {
	counter := 0
	volume := content
	edge := (float64(threshold) * content) / 100
	for volume >= edge {
		volume = volume - (volume * (float64(evapPerDay) / 100))
		counter++
	}
	return counter
}

func main() {
	fmt.Println(Evaporator(100, 5, 5))

}
