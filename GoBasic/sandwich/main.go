// 절차적 프로그래밍
// 절차적 프로그래밍은 유지보수, 수정이 어려움
// 가령 아래의 간단한 예에서 jam을 orange jam으로 바꾸려고 하면 바꾸어야 하는 코드가 많은데, 실제 실무에서는 유지보수의 cost가 매우 커진다.
// 이러한 문제 때문에 OOP가 등장하게 되었다.

package main

import "fmt"

type Bread struct {
	val string
}

type StrawberryJam struct {
	opened bool
}

type SpoonOfStrawberry struct {
}

type Sandwitch struct {
	val string
}

func GetBreads(num int) []*Bread {
	breads := make([]*Bread, num)
	for i := 0; i < num; i++ {
		breads[i] = &Bread{val: "bread"}
	}
	return breads
}

func OpenStrawberryJam(jam *StrawberryJam) {
	jam.opened = true
}

func GetOneSpoon(_ *StrawberryJam) *SpoonOfStrawberry {
	return &SpoonOfStrawberry{}
}

func PutJamOnBread(bread *Bread, jam *SpoonOfStrawberry) {
	bread.val += " + Strawberry Jam"
}

func MakeSandwitch(breads []*Bread) *Sandwitch {
	sandwitch := &Sandwitch{}
	for i := 0; i < len(breads); i++ {
		sandwitch.val += breads[i].val + " + "
	}
	return sandwitch
}

func main() {
	// 1. 빵 두개를 꺼낸다.
	breads := GetBreads(2)

	jam := &StrawberryJam{}

	// 2. 딸기잼 뚜껑을 연다.
	OpenStrawberryJam(jam)

	// 3. 딸기잼을 한스푼 뜬다.
	spoon := GetOneSpoon(jam)

	// 4. 딸기잼을 빵에 바른다.
	PutJamOnBread(breads[0], spoon)

	// 5. 빵을 덮는다.
	sandwitch := MakeSandwitch(breads)

	// 6. 완성
	fmt.Println(sandwitch.val)
}
