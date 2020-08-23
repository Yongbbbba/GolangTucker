package main

import "fmt"

func main() {
	var i int
	for i = 0; i < 10; i++ { //i의 scope는 for 문 안까지만, 그렇기 때문에 i를 for문 밖에서 선언해주고 값을 할당해줘야함
		fmt.Println(i)
	}
	fmt.Println("최종 i 값 :", i)
}

/*

i := 0
for i < 10 {
	fmt.Println(i)
	i++
}

위의 for문과 같은 의미임

*/
