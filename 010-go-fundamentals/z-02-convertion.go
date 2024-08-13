package main

import "fmt"

func convertion() {
	fmt.Println("convertion")
	var numberA int = 5
	fmt.Println(numberA)
	var numberB float32 = float32(numberA)
	fmt.Println(numberB)
	var numberC float64 = float64(numberB)
	fmt.Println(numberC)
	fmt.Println("convertion")
}