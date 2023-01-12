package main

import (
	"fmt"
)

func main() {
	//everything that we do with arrays we can do with slices except 1 or 2 exceptions
	// fmt.Println("Hello slices!!")
	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// b := a[:]   //slice all elements
	// c := a[3:]  //slice from 4th element
	// d := a[:6]  //slice first 6 elements
	// e := a[3:6] //slice 4th, 5th and 6th elements
	// a[5] = 43 //this will reflect in all arrays
	// b[0] = 2 //this will reflect in all arrays
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)

	//Built in make() function
	// a := make([]int, 3, 100) //length is 3 and capacity is 100
	// fmt.Println(a)
	// fmt.Println("Length: %v\n", len(a))
	// fmt.Println("Capacity: %v\n", cap(a))

	// a := []int{}
	// fmt.Println(a)
	// fmt.Printf("length= %v\n", len(a))
	// fmt.Printf("capacity= %v\n", cap(a))
	// a = append(a, 1, 2)
	// fmt.Println(a)                       //[1 2]
	// fmt.Printf("length= %v\n", len(a))   //2
	// fmt.Printf("capacity= %v\n", cap(a)) //2
	// a = append(a, []int{3, 4, 5}...)     //concanecate to previous values
	// fmt.Println(a)                       //[1,2,3,4,5]
	// fmt.Printf("length= %v\n", len(a))   //2
	// fmt.Printf("capacity= %v\n", cap(a)) //2

	a := []int{1, 2, 3, 4, 5}
	// b := a[:len(a)-1]

	// fmt.Println(b) //[1 2 3 4]

	// b := append(a[:2])
	// fmt.Println(b)//[1 2]
	// fmt.Println(a[4]) //5
	b := append(a[:2], a[3:]...) //appending elements from index 0 to 2 and 3 to 5
	fmt.Println(b)               //[1 2 4 5]

}
