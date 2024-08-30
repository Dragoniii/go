package main

import (
	"fmt"
)

func ar0() {
	// var arr0 [3]int
	// fmt.Println(arr0)
	// var arr1 [3]string
	// fmt.Println(arr1)
    // arr2 := [3]int{1, 2, 3}
	// fmt.Println(arr2)
    // fmt.Println(arr1[1])
    // arr1[1] = "a"
    // fmt.Println(arr1[1])
}

func ar1() {
ar0 := [3]string{"1a", "2b", "3c"}
fmt.Println(ar0)
ar1 := ar0
ok := ar0 == ar1
fmt.Println(ok)
fmt.Println(ar1)
fmt.Println("000")
ar0[0] = " 0"  
ar1[2] = "??"  
fmt.Println(ar0)
fmt.Println(ar1)
fmt.Println(len(ar0))
ok = ar0 == ar1
fmt.Println(ok)
}