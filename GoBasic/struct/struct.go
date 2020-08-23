package main

import "fmt"

//structure의 기본적인 개념 
// 일종의 객체라고 보면 됨 
// 속성(statement) 와 기능(method)로 구성

/*
type Person struct {
	name string
	age int
}

//Person 객체에 기능 추가하기 
func (p Person) PrintName(){
	fmt.Print(p.name)
}


func main()  {
	var p Person
	p1 := Person{"개똥이", 15}
	p2 := Person{name:"소똥이", age:21}
	p3 := Person{name:"Carson"}
	p4 := Person{}

	fmt.Println(p, p1, p2, p3, p4,)

	p.name = "Smith"
	p.age = 24

	fmt.Println(p)

	p.PrintName()
	
}

*/

type Student struct {
	name string
	class int
	sungjuk Sungjuk
}

type Sungjuk struct {
	name string
	grade string
}

func (s Student) ViewSungjuk()  {
	fmt.Println(s.sungjuk)
	
}

// func ViewSungjuk(s Student) {
// 	fmt.Println(s.sungjuk)
// }

func(s Student) InputSungjuk(name string, grade string) {
	s.sungjuk.name = name
	s.sungjuk.grade = grade
}

// func InputSungjuk(s Student, name string, grade string) {
// 	s.sungjuk.name = name
// 	s.sungjuk.grade = grade
// }

func main()  {

	var cheolsoo Student
	cheolsoo.name = "철수"
	cheolsoo.class = 1

	cheolsoo.sungjuk.name = "수학"
	cheolsoo.sungjuk.grade = "C"

	cheolsoo.ViewSungjuk()
	
	
	cheolsoo.InputSungjuk("과학", "A")
	cheolsoo.ViewSungjuk() // 왜 {과학,A} 가 아닌 {수학, C} 가 출력되는지 이해하지 못함 
}