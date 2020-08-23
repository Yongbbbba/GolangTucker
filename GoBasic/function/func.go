// fucntion practice
// 모든 재귀 호출은 반복문으로 바꿀 수 있음
// 상황에 따라 재귀 호출 또는 반복문이 쉬운 경우가 존재
// 재귀호출은 피보나치 수열에 유리
// 수학적 정의의 경우에는 재귀호출로 정의하는 것이 더 쉬운 경우가 많다
// 함수호출과정이 반복문보다 느리다, 메모리도 더 많이 사용함

package main

import "fmt"

func main() {
	s := sum(10, 0)
	fmt.Println(s)
}

func sum(x int, s int) int {
	if x == 0 {
		return s
	}
	s += x // s = s + x
	return sum(x-1, s)
}
