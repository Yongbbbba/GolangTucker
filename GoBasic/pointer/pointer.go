//포인터 개념 이해

package main

import "fmt"

// func main() {
// 	var a int
// 	var p *int
// 	var b int

// 	a = 3
// 	p = &a
// 	b = 2
// 	fmt.Println(a)
// 	fmt.Println(p)
// 	fmt.Println(*p)

// 	p = &b

// 	fmt.Println(b)
// 	fmt.Println(p)
// 	fmt.Println(*p)

// }

// 저번 시간 구조체 문제
type Student struct {
	name string
	age  int

	grade string
	class string
}

func (s *Student) PrintSungjuk() { // (s Student) 부분이 python class에서 self의 기능을 하는 것이 아닌가 싶다.
	fmt.Println(s.class, s.grade)
}

func (s *Student) InputSungjuk(class string, grade string) { //포인터 형태로 만들면 method의 역할을 하게끔 만든다
	s.class = class
	s.grade = grade

}

func main() {
	var s Student = Student{name: "Tucker", age: 23, class: "수학", grade: "A+"}
	s.InputSungjuk("과학", "C")
	s.PrintSungjuk()
}

// func main() {

// 	var a int
// 	a = 1
// 	Increase(&a)
// 	fmt.Println(a)

// }

// func Increase(x *int) {
// 	// x++
// 	*x = *x + 1

// }
