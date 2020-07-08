// array, string 연습
// 문자열 => 배열

package main

import "fmt"

// 배열 복사하기
/*
func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	clone := [5]int{}

	for i := 0; i < 5; i++ {
		clone[i] = arr[i]
	}

	fmt.Println(clone)
}
*/

// 배열 순서 바꾸기(복사의 방법 사용)
/*
func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	temp := [5]int{}
	for i := 0; i < len(arr); i++ {
		temp[i] = arr[len(arr)-1-i]
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = temp[i]
	}
	fmt.Println("temp: ", temp)
	fmt.Println("arr: ", arr)
}
*/

// 배열 순서 바꾸기(그냥 바꾸기)

/*
func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]

	}
	fmt.Println(arr)
}
*/

// Radix Sort -> 모든 경우가 아닌 일부 경우에 사용할 수 있는 정렬
// 조건만 맞으면 가장 정렬이 빠른 알고리즘
// 배열에 들어가는 원소의 갯수가 한정되어있고(실수의 경우 불가), 범위가 작야아지 사용가능(메모리 사용을 많이 함)

func main() {
	arr := [11]int{0, 5, 4, 9, 1, 2, 8, 3, 6, 4, 5}
	temp := [10]int{} //초기값이 0으로 되어 있음

	for i := 0; i < len(arr); i++ {
		idx := arr[i]
		temp[idx]++
	}
	fmt.Println("temp : ", temp)
	idx := 0
	for i := 0; i < len(temp); i++ {
		for j := 0; j < temp[i]; j++ {
			arr[idx] = i
			idx++
		}
	}

	fmt.Println("arr : ", arr)
}
