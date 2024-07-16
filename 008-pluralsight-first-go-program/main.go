// package main

// import "fmt"

// func main () {
// 	fmt.Println("Bora")
//}
//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//

// package main

// import "fmt"

// func main() {
// 	var a string
// 	a = "foo"
// 	fmt.Println(a)

// 	var b int = 99
// 	fmt.Println(b)

// 	var c = false
// 	fmt.Println(c)

// 	var d = 3.1415
// 	fmt.Println(d)

// 	var f float32 = float32(d)
// 	fmt.Println(f)

// 	myName := "Joel"
// 	fmt.Println(myName)

// }
// -//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//
// package main

// import "fmt"

// func main() {
// 	a, b := 5, 2

// 	fmt.Println(a + b)
// 	fmt.Println(a - b)
// 	fmt.Println(a * b)
// 	fmt.Println(a / b)
// 	fmt.Println(a % b)

// 	fmt.Println(float32(a) / float32(b))

//		fmt.Println(a == b)
//		fmt.Println(a > b)
//		fmt.Println(a < b)
//	}
//
// -//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//
// package main

// import "fmt"

// func main() {
// 	const a = 13
// 	var i int = a
// 	var j float64 = a

// 	fmt.Println(i, j)

// 	const c = iota
// 	fmt.Println(c)

// 	const (
// 		d = 2 * 5
// 		e
// 		f = iota
// 		g
// 		h = 10 * iota
// 	)
// 	fmt.Println(d, e, f, g, h)

// }
// -//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//
package main

import "fmt"

func main() {
	s := "Hello"

	p := &s

	fmt.Println(p)
	fmt.Println(*p)

	*p = "Jojo"

	fmt.Println(s)

}
