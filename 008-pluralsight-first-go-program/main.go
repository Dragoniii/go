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
// package main

// import "fmt"

// func main() {
// 	s := "Hello"

// 	p := &s

// 	fmt.Println(p)
// 	fmt.Println(*p)

// 	*p = "Jojo"

// 	fmt.Println(s)

// }
// -//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//
// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Go is great at making web services and applications!") })

// 	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "./home.html") })

//		http.ListenAndServe(":3000", nil)
//	}
//
// -//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main () {
	level := flag.String("level", "CRITICAL", "log level to filter for")
	flag.Parse()

	f, err := os.Open("./log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bufReader := bufio.NewReader(f)

	for line, err :=  bufReader.ReadString('\n'); err == nil; line, err = bufReader.ReadString('\n') {
		if strings.Contains(line, *level) {
			fmt.Println(line)
		}
		
	}
}