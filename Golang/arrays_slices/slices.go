package main

import (
	"fmt"
)

func main() {
	//everything that we do with arrays we can do with slices except 1 or 2 exceptions
	// fmt.Println("Hello slices!!")
	a := []int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

}
