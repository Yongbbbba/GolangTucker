package main

import (
	"fmt"
)

type StructA struct {
	val string
}

func (s *StructA) String() string {
	return "Val:" + s.val

}

func main() {
	a := &StructA{val: "AAA"}
	fmt.Println(a)

}
