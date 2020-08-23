// interface와 type을 이용한 golang에서의 OOP
// object는 상태와 기능을 가지고 있음
// interface는 object에서 기능만을 구현하고 관계를 따로 정의한(decoupling) type이다.
// 이렇게함으로써 종속성을 없애고 확장을 용이하게 한다.
// 여기서 기능은 공개기능(외부기능)과 내부기능이 있는에 공개기능이 관계라고 볼 수 있다.
// 절차형 프로그래밍과 다르게 새로운 jam(객체)을 추가하기가 용이하다.

package main

import (
	"fmt"
)

type SpoonOfJam interface {
	String() string // 이름(입력) 출력 -> 인터페이스 선언
}

type Jam interface {
	GetOneSpoon() SpoonOfJam
}

type Bread struct {
	val string
}

func (b *Bread) PutJam(jam Jam) {
	spoon := jam.GetOneSpoon() // 한 숟갈 떠서
	b.val += spoon.String()    // 빵에다 바르기, "bread + jam"
}

func (b *Bread) String() string {
	return "bread " + b.val
}

type StrawberryJam struct {
}

func (j *StrawberryJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfStrawberryJam{}
}

type OrangeJam struct {
}

func (j *OrangeJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfOrangeJam{}
}

type AppleJam struct {
}

func (j *AppleJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfAppleJam{}
}

type SpoonOfStrawberryJam struct {
}

func (s *SpoonOfStrawberryJam) String() string {
	return "+ strawberry"
}

type SpoonOfOrangeJam struct {
}

func (s *SpoonOfOrangeJam) String() string {
	return "+ Orange"
}

type SpoonOfAppleJam struct {
}

func (s *SpoonOfAppleJam) String() string {
	return "+ Apple"
}

func main() {
	bread := &Bread{}
	//jam := &StrawberryJam{}
	//jam := &OrangeJam{}
	jam := &AppleJam{}
	bread.PutJam(jam)

	fmt.Println(bread)
}
