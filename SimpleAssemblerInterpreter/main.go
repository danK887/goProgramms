package main

import (
	"fmt"
	"strconv"
	"strings"
)

func SimpleAssembler(program []string) map[string]int {
	value := make(map[string]int)
	instrDecomp := []string{}
	for i := 0; i < len(program); i++ {
		instruction := program[i]
		instrDecomp = strings.Fields(instruction)
		for _, val := range instrDecomp {
			switch val {
			case "mov":
				value[instrDecomp[1]], _ = strconv.Atoi(instrDecomp[2])
			case "inc":
				value[instrDecomp[1]]++
			case "dec":
				value[instrDecomp[1]]--
				// case "jnz":
				// 	if value[instrDecomp[1]] == 0 {
				// 		break
				// 	} else if instrDecomp[2] == "-1" {

				// 	}
			}
			break
		}
	}

	fmt.Println(instrDecomp)
	return value
}

func main() {
	fmt.Println(SimpleAssembler([]string{"mov a 1", "inc a", "dec a", "dec a", "jnz a -1", "inc a"})) //{"a": 1}
}
