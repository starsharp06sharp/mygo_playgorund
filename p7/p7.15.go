package main

import (
	"bytes"
	"fmt"
)

func uniStr(s string) string {
	rs := []rune(s)
	var buf bytes.Buffer
	var lastRune rune = 0
	for _, r := range rs {
		if r != lastRune {
			lastRune = r
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

func main() {
	var s string
	fmt.Scanln(&s)
	fmt.Println(uniStr(s))
}
