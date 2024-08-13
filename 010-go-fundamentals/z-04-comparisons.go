package main

import "fmt"

func comparisons() {
	fmt.Println("comparisons")
	numberA, numberB := 3, 3
	numberC := numberA == numberB
	fmt.Println(numberA, numberB, numberC)
	// numberD := -3
	// numberE := -3.0	
	// numberF := numberD == numberE
	// fmt.Println(numberD, numberE, numberF)
	// Golang does not accept this comparison
	fmt.Println("comparisons")
}
