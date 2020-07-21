// channel 학습
// channel은 fixed sized, thread safe한 queue다.
// channel도 일종의 type, 자료구조 중 하나이다. (chan)
// var a chan int
// a := make(chan int, 10)
// a:= make(chan int) -> 0개 짜리 channel임.
// queue이기 때문에 push와 pop으로 기능이 이루어짐

package main

import (
	"fmt"
)

func pop(c chan int) {
	fmt.Println("Pop func")
	v := <-c
	fmt.Println(v)

}

func main() {
	var c chan int
	c = make(chan int)

	// c <- 10 // 0개 짜리 channel을 만들 경우에 다른 곳에서 이것을 빼주기 전까지는 deadlock에 걸림
	go pop(c)
	c <- 10

	fmt.Println("end of program")

}
