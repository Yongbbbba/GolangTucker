package main

import (
	"fmt"

	"github.com/yongbbbba/tucker/GoBasic/dataStruct"
)

func main() {

	fmt.Println("abcde = ", dataStruct.Hash("abcde"))
	fmt.Println("abcdf = ", dataStruct.Hash("abcdf"))
	fmt.Println("tbcde = ", dataStruct.Hash("tbcde"))
	fmt.Println("zbcde = ", dataStruct.Hash("zbcde"))

	m := dataStruct.CreateMap()
	m.Add("AAA", "01077777777")
	m.Add("BBB", "01088888888")
	m.Add("CDEAFDGEAGFD", "011111111")
	m.Add("CCC", "0102329023")

	fmt.Println("AAA = ", m.Get("AAA"))
	fmt.Println("CCC = ", m.Get("CCC"))
	fmt.Println("CDEAFDGEAGFD = ", m.Get("CDEAFDGEAGFD"))
	fmt.Println("DDD = ", m.Get("DDD"))

}
