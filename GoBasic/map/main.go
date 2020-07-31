package main

import "fmt"

func main() {
	var m map[string]string // map 선언
	m = make(map[string]string)

	m["bcd"] = "ccc"
	m["abc"] = "bbb"

	fmt.Println(m["bcd"])

	m1 := make(map[int]string)
	m1[53] = "ddd"
	fmt.Println(m1[55])

	m2 := make(map[int]int)
	m2[4] = 4
	m2[5] = 0
	m2[2] = 2
	m2[10] = 10

	for key, value := range m2 {
		fmt.Println(key, " = ", value)
	}

}
