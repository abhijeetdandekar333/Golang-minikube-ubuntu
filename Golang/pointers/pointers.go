package main

import "fmt"

func main() {
	// var a int = 42
	// var b *int = &a
	// a = 5
	// fmt.Println(a, b, *b)
	// *b = 14
	// fmt.Println(a, b, *b)

	var ms *myStruct
	//Printing values with &
	ms = &myStruct{foo: 42}
	fmt.Println(ms) //&{42}
	// OR
	ms = new(myStruct)
	fmt.Println(ms) //&{0}

	(*ms).foo = 42
	fmt.Println((*ms).foo) //42
	// OR
	ms.foo = 53
	fmt.Println((ms.foo)) //53

}

type myStruct struct {
	foo int
}
