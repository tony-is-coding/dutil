package main

import "fmt"

func main() {

	var a int32

	a = 1
	fmt.Println(a << 34)

	k := 33
	const n = 32
	s := k & (n - 1)
	fmt.Println(a<<s | a>>(n-s))
}
