package main

import "fmt"

func main() {
	// fmt.Println("hello maps")
	//map contains key value pair
	// syntax -> map[key-datatype]value-datatype

	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"india":     91,
		"america":   1,
		"australia": 21,
		"england":   22,
		"france":    23,
		"argentina": 24,
		"japan":     25,
	}
	_, ok := statePopulations["oho"]
	fmt.Println(ok)
	statePopulations["pakistan"] = 26 //to add entry to map
	fmt.Println(statePopulations["india"])
	delete(statePopulations, "pakistan") //to delete entry from a map
	fmt.Println(statePopulations)
	fmt.Println(len(statePopulations))
}
