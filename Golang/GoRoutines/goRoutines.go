package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello("Hello")
	sayHello("World!!")
}

func sayHello(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}
