package main

import "fmt"

type Son struct {
	Name string
	Age  int
}

type Implement struct {
	Son
	High uint
}

func main() {

	//var a int32
	//
	//a = 1
	//fmt.Println(a << 34)
	//
	// 32 位实现循环移位
	//k := 33
	//const n = 32
	//s := k & (n - 1)
	//fmt.Println(a<<s | a>>(n-s))
	i := new(Implement)
	fmt.Println(i.Son.Age)


}
