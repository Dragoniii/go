package main

import (
	"fmt"
)

func s0() {
	fmt.Println("??????????????????")
	s0 := []string{"1a", "2b", "3c"}
	s1 := s0
	fmt.Println(s1)
	fmt.Println(s0)
	fmt.Println("slice")
	s0[1] = " 0"  
	fmt.Println(s0)
	fmt.Println(s1)
	fmt.Println("slice")
	s0 = append(s0, "1")
	s1 = append(s1, "0")
	fmt.Println(s0)
	fmt.Println(s1)
	fmt.Println("slice")
	fmt.Println(len(s0))
	
	s0[0] = " 0"  
	fmt.Println(s0)
	fmt.Println(s1)
}
