// 피보나치 수열 만들기
// 1 1 2 3 5 8 13 21 34 55

package main

import (
	"fmt"
	"time"
)

func main() {

	// 두 알고리즘의 실행시간 비교
	// 재귀호출의 경우
	startTime := time.Now()
	result := Fibonacci1(50)
	runTime := time.Since(startTime)
	fmt.Printf("재귀호출을 이용한 피보나치 수열 알고리즘 결과는 %d입니다\n", result)
	fmt.Printf("실행시간은 %s입니다\n", runTime)
	startTime = time.Now()
	result = Fibonacci2(50)
	runTime = time.Since(startTime)
	fmt.Printf("반복문을 이용한 피보나치 수열 알고리즘 결과는 %d입니다\n", result)
	fmt.Printf("실행시간은 %s입니다\n", runTime)
}

// 재귀호출을 이용한 피보나치 수열
func Fibonacci1(x int) int {

	if x <= 2 {
		return 1
	} else {
		return Fibonacci1(x-1) + Fibonacci1(x-2)
	}
}

// 재귀호출이 아닌 반복문을 통한 피보나치 수열
func Fibonacci2(n int) int {
	var index int
	var last1 int
	var last2 int
	var result int

	if n <= 2 {
		return 1
	}
	last1 = 1
	last2 = 1
	for index = 2; index < n; index++ {
		result = last1 + last2
		last1 = last2
		last2 = result
	}
	return result
}
