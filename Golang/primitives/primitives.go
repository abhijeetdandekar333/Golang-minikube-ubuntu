package main

import "fmt"

func main() {
	// var n bool = true
	// fmt.Println(n)
	// var o bool
	// n := 1 == 2
	// m := 2 == 2
	// fmt.Printf("%v, %T\n", n, n)
	// fmt.Printf("%v, %T\n", m, m)
	// fmt.Printf("%v, %T", o, o) //false

	var n uint16 = 42
	fmt.Printf("%v, %T\n", n, n)

	a := 10             //1010
	b := 3              // 0011
	fmt.Println(a & b)  // 0010
	fmt.Println(a | b)  //1011
	fmt.Println(a ^ b)  // 1001
	fmt.Println(a &^ b) // 0100

	//Bit shifting
	c := 8              //2^3
	fmt.Println(c << 3) //2^3 * 2^3 = 2^6
	fmt.Println(c >> 3) // s^3 / 2^3 = 2^0

	//complex numbers
	var m complex64 = complex(5, 13)
	fmt.Printf("%v, %T\n", m, m)

	//strings
	// s := "This is a string"
	// fmt.Printf("%v , %T\n", string(s[2]), s[2])
	// q := "This is a string 2"
	// fmt.Printf("%v , %T\n", s+q, s+q)
	s := "This is a string"
	q := []byte(s)
	fmt.Printf("%v, %T\n", string(q), string(q))

	// rune ->
	// r := 'a'
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)
}
