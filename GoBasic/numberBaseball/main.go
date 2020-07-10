package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	//  무작위로 숫자 3개를 만든다.
	numbers := MakeNumbers()

	cnt := 0
	for {
		cnt++
		// 사용자의 입력을 받는다.
		inputNumbers := InputNumbers()

		// 결과를 비교한다.
		result := CompareNumbers(numbers, inputNumbers)

		PrintResult(result)

		// 3S 인가 ?
		if IsGameEnd(result) {
			// 게임 끝
			break
		}
	}

	// 게임끝. 몇 번만에 맞췄는지 출력
	fmt.Printf("%d 번만에 맞췄습니다.\n", cnt)
}

func MakeNumbers() [3]int {
	// 0~9 사이의 겹치지 않는 무작위 숫자 3개를 반환한다.
	var rst [3]int
	for i := 0; i < 3; i++ {
		for { //중복되지 않은 숫자를 뽑을 때가지 무한루프

			n := rand.Intn(10) // 0~9까지 숫자 중 랜덤으로 하나 뽑기
			duplicated := false
			for j := 0; j < i; j++ { //자기 순서보다 앞에 것까지 중복을 확인
				if rst[j] == n {
					// 겹친다. 다시 뽑는다.
					duplicated = true
					break
				}
			}
			if !duplicated { //duplicated가 false라면
				rst[i] = n
				break //무한루프 빠져나가기
			}
		}

	}
	fmt.Println(rst)
	return rst
}

func InputNumbers() [3]int {
	// 키보드로부터 0~9 사이의 겹치지 않는 숫자 3개를 입력받아 반환한다.
	var rst [3]int
	return rst
}

func CompareNumbers(numbers, inputNumbers [3]int) bool {
	// 두 개의 숫자 3개를 비교해서 결과를 반환한다.
	return true
}

func PrintResult(result bool) {
	fmt.Println(result)

}

func IsGameEnd(result bool) bool {
	// 비교 결과가 3 스트라이크인지 확인

	return result

}
