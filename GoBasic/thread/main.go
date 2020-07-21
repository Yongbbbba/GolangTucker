//thread 학습
//go thread는 context swithing을 프로그래머가 신경쓰지 않게 도와주는 역할(?)
//cpu 갯수보다 thread가 많게 되면 context swithing이 일어나게 되고, context swithing은 비용을 수반한다.
//그렇기 때문에 context switch을 최대한 적게, 효율적으로 하는 것이 중요하다.
// 메모리에 적재되어 실행되고 있는 프로그램을 프로세스라고 한다.
// 프로세스는 여러 개의 쓰레드를 가질 수 있음.
// 프로그램은 꼭 하나의 실행파일로만 이루어져있는 것은 아니다.
// 프로그램은 논리적인 개념으로, 기능을 묶어놓은 것이라고 보면 됨.
// deadlock은 lock이 막혔다는 것

package main

import (
	"fmt"
	"time"
)

func main() {
	fun1() // go fun1(), go thread를 만들어서 func1을 수행하라는 의미, func1의 쓰레드와 main의 쓰레드는 달라지게 됨
	for i := 0; i < 20; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("main", i)
	}
	fmt.Scanln()

}

func fun1() {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("func1: ", i)
	}
}
