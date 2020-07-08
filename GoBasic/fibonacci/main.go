// 피보나치 수열 재귀호출을 통해 만들기

package main

import "fmt"

func main() {
	result := f(10)
	fmt.Println(result)
}

func f(x int) int {
	if x == 0 {
		return 1
	}
	if x == 1 {
		return 1
	}
	return f(x-1) + f(x-2)
}
