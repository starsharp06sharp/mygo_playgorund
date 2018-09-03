package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var feb = func() func() int {
		a, b := 1, 1
		return func() int {
			a, b = b, a+b
			return a
		}
	}()
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", feb())
	}
}
