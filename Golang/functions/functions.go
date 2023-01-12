package main

import "fmt"

// func main() {
// 	sum("The sum is : ", 1, 2, 3, 4, 5)
// }
// func sum(msg string, values ...int) {
// 	fmt.Println(values)
// 	result := 0
// 	for _, v := range values {
// 		result += v
// 	}
// 	fmt.Println(msg, result)
// }

// func main() {
// 	d, err := divide(5.0, 0.0)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(d)
// }

// func divide(a, b float64) (float64, error) {
// 	if b == 0.0 {
// 		return 0.0, fmt.Errorf("Cannot divide by zero")
// 	}
// 	return a / b, nil
// }

func main() {
	// for i := 0; i < 5; i++ {
	// 	func(i int) {
	// 		fmt.Println(i)
	// 	}(i)
	// }

	var f func() = func() {
		fmt.Println("Hello Go")
	}
	f()

}
