// slice leture
// slice 는 3개의 properties를 가지고 있음
// pointer(배열의 시작주소, len, cap)
// append를 할 때, capcacity에 여유가 있으면 그 address의 여유공간에 추가를 하지만, capacity가 부족할 경우 새로운 공간을 만들어 새로운 메모리 주소를 가리킨다. 그리고 보통 capacity는 2배로 늘린다.
// object는 property와 method를 가지고 있다.
// object가 되면 subject와 verb를 가지게 되었다. 주어(주체가)가 나타나기 시작함. oop는 주어와 오브젝트의 관계를 이야기한다.
// 기존 절자지향에서는 기능이 중요하였지만, oop에서는 object와 object간의 relation이 중요해짐.
// 이런 subject가 instance임. 하나의 생명주기를 가지게 됨.
// 표현 방식이 달라질 뿐 컴퓨터 입장에서는 똑같음.
// 쉽게 말하면 instance란 포인터 타입의 structure라고 이해해도 무방함.

package main

import "fmt"

type Student struct {
	name  string
	age   int
	grade int
}

func (t *Student) SetName(newName string) {
	t.name = newName
}

func main() {
	a := Student{"aaa", 20, 10}
	fmt.Println(a)
	a.SetName("bbb")
	fmt.Println(a)
}
