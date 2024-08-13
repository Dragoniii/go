package main

import "fmt"

func printHelloWorld() {
	fmt.Println("---")
	fmt.Println(" 01 - Hello, \n World!")
	fmt.Println(" 01 - Hello, World!")
	fmt.Println(
	`01 - 
	Hello, 
	World!`)
	fmt.Println(` 01 - Hello, \n World!`)
	fmt.Println("---")
}
