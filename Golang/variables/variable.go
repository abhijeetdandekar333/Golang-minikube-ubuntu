package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var i int
	// i = 42
	// // j := 27
	// // var j int = 27
	// fmt.Println(i)
	// // fmt.Println(j)
	// // j := 99.99
	// fmt.Printf("%v, %T", i, i)

	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j float32

	j = float32(i) //assigns value of i to j in float format
	fmt.Printf("%v, %T\n", j, j)

	var k string
	k = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", k, k)

}
