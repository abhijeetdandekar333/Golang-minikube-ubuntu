package main

import (
	"fmt"
	// "reflect"
)

type Animal struct {
	Name   string //string can have max of 100 characters
	Origin string
}

type Bird struct {
	Animal   //embedding animal struct into the bird struct
	SpeedKPH float32
	CanFly   bool
}

func main() {
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)

	//Literal syntax
	c := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"}, //need to define Animal: Animal{}
		SpeedKPH: 78,
		CanFly:   false,
	}
	fmt.Println(c)

}
