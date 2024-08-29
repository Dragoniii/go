package main

import "fmt"

func main() {
	// 01
	handler()
	fmt.Println("Item 1 has been skipped")
	// 02
	marshall()
	// 03
	unmarshall()
	// 04
	middleware()
}
