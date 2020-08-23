package main

import "fmt"

func main() {
	a := 4
	if a == 3 {
		fmt.Println("a = 3")

	} else if a == 4 {
		fmt.Println("a = 4")
	} else {
		fmt.Println("a는 3도 아니고 4도 아니다")
	}
}
