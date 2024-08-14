package main

import "fmt"

func varandconst() {
	fmt.Println("varandconst")
	const numberA = 42
	const numberB float32 = numberA 
	const numberC float32 = numberB
	const numberD float64 = float64(numberC) 
	fmt.Println(numberA, numberB, numberC, numberD)
	fmt.Println("---")
	const a = iota
	const b = iota
	const c = 5 * 2
	const d = iota
	const (
		aa = 0
		bb = 5 * 2
		cc
		dd = iota
		ee
	)
	fmt.Println(a, b, c, d)
	fmt.Println(aa, bb, cc, dd, ee)
	fmt.Println("varandconst")
}