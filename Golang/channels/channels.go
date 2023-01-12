package main

import (
	"fmt"
	"sync"
)

func main() {
	// fmt.Println("Hello Channels")

	myCh := make(chan int, 3) //defining channel using make

	wg := &sync.WaitGroup{}
	// fmt.Println(<-myCh)
	// myCh <- 5 //pushing value into channel
	wg.Add(2)
	//receive only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myCh

		fmt.Println(isChannelOpen)
		fmt.Println(val)
		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)
	//send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 5
		myCh <- 6
		myCh <- 0
		close(myCh)

		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
