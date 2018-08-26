package main

import (
	"fmt"
)

func splitStr(s string, i int) (string, string) {
	return s[:i], s[i:]
}

func main() {
	fmt.Println(splitStr("abc1234", 3))
}
