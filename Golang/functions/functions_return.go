package main

import "fmt"

// func main() {
// 	s := sum(1, 2, 3, 4, 5)
// 	fmt.Println("The sume is ", s)
// }
// func sum(values ...int) int { //int is data-type of return value
// 	fmt.Println(values)
// 	result := 0
// 	for w, v := range values { //w -> index , v -> values
// 		result += v
// 		fmt.Println(w) //prints indexes
// 	}
// 	return result //returns value of result
// }

// By OR using pointers

// func main() {
// 	s := sum(1, 2, 3, 4, 5)
// 	fmt.Println("The sume is ", *s)
// }
// func sum(values ...int) *int { //int is data-type of return value
// 	fmt.Println(values)
// 	result := 0
// 	for w, v := range values { //w -> index , v -> values
// 		result += v
// 		fmt.Println(w) //prints indexes
// 	}
// 	return &result //returns value of result
// }

// OR

func main() {
	s := sum(1, 2, 3, 4, 5)
	fmt.Println("The sume is ", s)
}
func sum(values ...int) (result int) { //int is data-type of return value
	fmt.Println(values)
	for _, v := range values { //w -> index , v -> values
		result += v
		// fmt.Println(w) //prints indexes
	}
	return //returns value of result
}
