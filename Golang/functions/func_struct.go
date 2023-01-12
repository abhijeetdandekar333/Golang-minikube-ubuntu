package main

import "fmt"

func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() { //greet() is a method
	fmt.Println(g.greeting, g.name) //(g greeter) gives contect of greeter
}

// (g greeter) creates copy of struct
//if we use (g *greeter) it changes the original struct
