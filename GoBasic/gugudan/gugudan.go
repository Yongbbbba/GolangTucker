package main

import "fmt"

/*

  *
 ***
*****
 ***
  *
마름모 만들기

*/

func main() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 2-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 3-2*i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
