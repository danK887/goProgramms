package main

import (
	"fmt"
)

type MyStruct struct {
	name string
	age  int
	sex  string
}

func userInfo(m *MyStruct) {
	m.name = "gge"
}

func main() {
	var m MyStruct
	userInfo(&m)
	fmt.Println(m.name)

	a := 12
	b := &a
	fmt.Println(*b)

}
