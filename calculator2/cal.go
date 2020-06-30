package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("숫자를 입력하세요")
	reader := bufio.NewReader(os.Stdin) // 표준 입력을 인자로 받아 읽기 객체를 만드는 것
	line, _ := reader.ReadString('\n')  // 한 줄 띄는 것 직전까지 string 읽기  한 줄과 에러를 읽는 것임
	line = strings.TrimSpace(line)      //빈 공간 날려버리기

	n1, _ := strconv.Atoi(line) //문자를 숫자로 바꾸는 것

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n2, _ := strconv.Atoi(line)

	fmt.Printf("입력하신 숫자는 %d, %d 입니다.", n1, n2)

	fmt.Println("연산자를 입력하세요")

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	if line == "+" {
		fmt.Printf("%d + %d = %d", n1, n2, n1+n2)
	} else if line == "-" {
		fmt.Printf("%d - %d = %d", n1, n2, n1-n2)
	} else if line == "/" {
		fmt.Printf("%d / %d = %d", n1, n2, n1/n2)
	} else if line == "*" {
		fmt.Printf("%d - %d = %d", n1, n2, n1*n2)
	} else {
		fmt.Println("잘못 입력하셨습니다.")
	}
}
