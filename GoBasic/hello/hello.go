package main

/*프로그램 시작점을 포함한 패키지는 main으로 시작해야함, 실행파일을 만들 때 그 실행파일이 어디서부터 시작하는지
프로그램 시작 포인트를 알려줘야함. 라이브러리는 기능을 묶어놓은 것이기 때문에 시작포인트를 명시할 필요가 없음 */

import "fmt" //go org에서 만든 go 표준 패키지임

// 함수 main도 특별한 함수, instruction point의 시작점을 의미 , 이 프로그램은 여기서부터 시작을 해 라고 알려 주는 것
func main() {
	fmt.Println("Hello world")
}
