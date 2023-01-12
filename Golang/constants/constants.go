package main

import (
	"fmt"
)

// package level
const myConst2 int = 22

// const (
//
//	b = iota
//	c = iota
//	d = iota
//
// )
// OR
const (
	b = iota
	c
	d
)
const (
	b1 = iota
)

func main() {
	//function level
	fmt.Println("Hello constants")

	const myConst int = 42
	fmt.Printf("%v,%T\n", myConst, myConst)
	const myConst1 bool = true
	fmt.Printf("%v,%T\n", myConst1, myConst1)
	// const myConst2 int = 42
	fmt.Printf("%v,%T\n", myConst2, myConst2)

	const a = 42.5 //compilers ability to detect datatype automatically
	fmt.Printf("%v,%T\n", a, a)

	//Enumerated constants
	fmt.Printf("%v,%T\n", b, b) //0
	fmt.Printf("%v,%T\n", c, c) //1
	fmt.Printf("%v,%T\n", d, d) //2
	//value pf iota automatically increases or increments

	fmt.Printf("%v,%T\n", b1, b1)
}
