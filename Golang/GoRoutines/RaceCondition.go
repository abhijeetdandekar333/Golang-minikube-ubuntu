package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition")

	wg := &sync.WaitGroup{} //used pointers
	mut := &sync.Mutex{}

	var score = []int{0}

	wg.Add(3) //number og go routines
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	wg.Wait()
	fmt.Println(score)
}
