package main

import (
	"fmt"
)

func reverseStr(s string) string {
	rrs := make([]rune, 0)
	for _, r := range []rune(s) {
		rrs = append([]rune{r}, rrs...)
	}
	return string(rrs)
}

func main() {
	var s string
	fmt.Scanln(&s)
	fmt.Println(reverseStr(s))
}
