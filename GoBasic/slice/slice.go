// slice leture

package main

import "fmt"

func RemoveBack(a []int) []int {
	return a[:len(a)-1]
}

func main() {
	a := []int{}
	for i := 0; i < 10; i++ {
		a[i] = i
	}
	fmt.Println(a)

	// for i := 0; i < 5; i++ {
	// 	a = RemoveBack(a)
	// }

	// fmt.Println(a)

}
