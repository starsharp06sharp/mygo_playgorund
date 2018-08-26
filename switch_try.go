package main

import "fmt"

func main() {
	a, b, c, d := 1, 2, 3, 2
	switch 2 {
	case a:
		fmt.Println("a is 2")
	case b:
		fmt.Println("b is 2")
	case c:
		fmt.Println("c is 2")
	case d:
		fmt.Println("d is 2")
	}
	// out: b is 2
	// failed :-(
}
