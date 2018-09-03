package main

import (
	"fmt"
	"strings"
)

func change(r rune) rune {
	if r > 255 {
		return rune('?')
	}
	return r
}

func main() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Println(strings.Map(change, s))
}
