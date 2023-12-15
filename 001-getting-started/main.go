package main

import (
	"fmt"

	"example.com/m/mascot"
	"rsc.io/quote"
)

func main () {
	fmt.Println(mascot.BestMascot())
	fmt.Println(quote.Go())
}