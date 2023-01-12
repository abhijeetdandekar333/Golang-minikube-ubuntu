package main

import (
	"fmt"
)

type Doctor struct {
	number     int
	actorname  string
	companions []string
}

func main() {

	aDoctor := Doctor{
		number:    3,
		actorname: "salman khan",
		companions: []string{
			"abhijeet", "sarvesh", "john",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorname)
	fmt.Println(aDoctor.companions[1])

	aDoctor1 := struct{ name string }{name: "ABhijeet"} //one line struct
	fmt.Println(aDoctor1)                               //{ABhijeet}
	aDoctor2 := &aDoctor1
	aDoctor2.name = "abhijeet"
	fmt.Println(aDoctor2) //&{abhijeet}

}
