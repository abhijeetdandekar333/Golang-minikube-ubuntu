package main

import (
	"fmt"
)

// package level
const (
	_               = iota //assigns 0 to unknown variable
	catSpecialist   = iota
	dogSpecialist   = iota
	snakeSpecialist = iota
)

func main() {
	//function level
	// fmt.Println("Hello constants1")
	var specialistType int
	fmt.Printf("%v\n", specialistType == catSpecialist)
	fmt.Printf("%v\n", catSpecialist)
}
