package main

import "fmt"

func pointer() {
	fmt.Println("pointer")
	a := 0
	fmt.Println(a)
	b := &a
	fmt.Println(a, b)
	c := b
	fmt.Println(a, b, *b, c, *c)
	*c = 6
	fmt.Println(a, b, *b, c, *c)
	fmt.Println("pointer")
}