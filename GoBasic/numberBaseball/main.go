// 숫자 야구 만들기

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	strikes int
	balls   int
}

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

			n := rand.Intn(10)       // 0~9까지 숫자 중 랜덤으로 하나 뽑기
			duplicated := false      // flag variable
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
	//fmt.Println(rst)
	return rst
}

func InputNumbers() [3]int {
	// 키보드로부터 0~9 사이의 겹치지 않는 숫자 3개를 입력받아 반환한다.
	var rst [3]int

	for {
		fmt.Println("겹치지 않는 0~9 사이의 숫자 3개를 입력해주세요.")
		var no int
		_, err := fmt.Scanf("%d\n", &no) // 입력을 몇 개 받았는지와 error가 return, Scanf는 숫자를 input으로 받음(?)
		// standard input을 받으면 키보드 버퍼에 숫자 3개 뿐만 아니고, 마지막에 엔터 키값도 함께 들어간다.
		// 그래서 만약에 중복되는 숫자를 넣어서 다시 루프가 돌아가면, 숫자 세 개는 버퍼에서 가져와서 사라지지만 엔터가 아직 남아있어서 그것을 뽑아오려고 하는데, 그것은 숫자가 아니라 에러가 난다.
		// 에러가 나면 바로 밑에 if 문에서 "잘못 입력하셨습니다."가 출력이 되기 때문에 버퍼에서 엔터도 없애기 위해 "%d\n" 이라고 쓴다.
		if err != nil { //에러가 있는 경우
			fmt.Println("잘못 입력하셨습니다.")
			continue
		}
		success := true
		idx := 0
		for no > 0 {
			n := no % 10 //나머지
			no = no / 10 // 몫

			duplicated := false // flag variable

			for j := 0; j < idx; j++ { //자기 순서보다 앞에 것까지 중복을 확인
				if rst[j] == n {
					// 겹친다. 다시 뽑는다.
					duplicated = true
					break
				}
			}

			if duplicated { //중복되는 숫자가 있다면
				fmt.Println("중복되는 숫자를 넣어서는 안됩니다.")
				success = false // flag variable
				break
			}

			if idx >= 3 { // 숫자가 3개를 초과하여 입력된 경우
				fmt.Println("3개보다 많은 숫자를 입력하셨습니다.")
				success = false
				break // 숫자 입력으로 다시 되돌아감
			}

			rst[idx] = n
			idx++ // 이런 식으로 쓰면 python의 while문과 좀 흡사한 것 같다
		}

		if success && idx < 3 { // 숫자가 3개 미만으로 입력된 경우
			fmt.Println("3개의 숫자를 입력하세요.")
			success = false
		}

		if !success {
			continue
		}

		break

	}

	rst[0], rst[2] = rst[2], rst[0]

	//fmt.Println(rst)
	return rst

}

func CompareNumbers(numbers, inputNumbers [3]int) Result {
	// 두 개의 숫자 3개를 비교해서 결과를 반환한다.
	strikes := 0
	balls := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			if numbers[i] == inputNumbers[j] {
				if i == j {

					strikes++
				} else {
					balls++
				}
				break // ????
			}
		}
	}

	return Result{strikes, balls}
}

func PrintResult(result Result) {
	fmt.Printf("%dS%dB\n", result.strikes, result.balls)

}

func IsGameEnd(result Result) bool {
	// 비교 결과가 3 스트라이크인지 확인

	return result.strikes == 3 // true 면

}
