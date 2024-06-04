package main

import "fmt"

func SimpleAssembler(program []string) map[string]int {
	// TODO: Implement your solution here
}

func main() {
	fmt.Println(SimpleAssembler([]string{"mov a 5", "inc a", "dec a", "dec a", "jnz a -1", "inc a"}))
}
