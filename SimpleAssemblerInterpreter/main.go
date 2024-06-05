package main

import (
	"fmt"
	"strconv"
	"strings"
)

func SimpleAssembler(program []string) map[string]int {
	value := make(map[string]int)
	instrDecomp := []string{}
	var err error
	for i := 0; i < len(program); i++ {
		instruction := program[i]
		instrDecomp = strings.Fields(instruction)
		for _, val := range instrDecomp {
			switch val {
			case "mov":
				value[instrDecomp[1]], err = strconv.Atoi(instrDecomp[2])
				if err != nil {
					value[instrDecomp[1]] = value[instrDecomp[2]]
				}
			case "inc":
				value[instrDecomp[1]]++
			case "dec":
				value[instrDecomp[1]]--
			case "jnz":
				instr, _ := strconv.Atoi(instrDecomp[2])
				if value[instrDecomp[1]] == 0 {
					break
				} else if instr < 0 {
					i = i + instr - 1
					break
				} else if instr > 0 {
					i = i + instr - 1
					break
				}
			}
			break
		}
	}
	return value
}

func main() {
	fmt.Println(SimpleAssembler([]string{"mov a -10", "mov b a", "inc a", "dec b", "jnz a -2"})) //{"a": 1}

}
