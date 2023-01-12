package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup //usually these are declared as pointers

var mut sync.Mutex //usually pointers to use for different goroutines

func main() {
	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1) //asks main to add the go routines currently only one
	}
	wg.Wait() //asks main to wait until all routines are added
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	defer wg.Done() //says the main function that the routines are added

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("oops in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
