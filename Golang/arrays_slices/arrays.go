package main

import (
	"fmt"
)

func main() {
	// grades := [3]int{1, 2, 3}
	// fmt.Printf("Grades: %v\n", grades)
	// fmt.Printf("Grades1: %v\n", grades[1])
	// fmt.Printf("Grades: %v\n", len(grades))

	//Array inside array 3d array
	// var identityMatrix [3][3]int = [3][3]int{ [3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	// fmt.Println(identityMatrix)

	a := [...]int{1, 2, 3}
	// b := a   //copied elements to array b
	b := &a         //b has address of a so any changes to b reflect in a also
	b[1] = 5        //changed second element of b to 5
	fmt.Println(a)  // for b := &a [1 5 3]
	fmt.Println(b)  // &[1 5 3]
	fmt.Println(*b) // [1 5 3]
}
