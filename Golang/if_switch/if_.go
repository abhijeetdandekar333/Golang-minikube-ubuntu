package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println("hello")
	if true {
		fmt.Println("The test is true")
	}
	myNum := 0.1
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}
}
