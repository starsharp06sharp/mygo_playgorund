package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Printf("byte count of s is:%d\n", len(s))
		fmt.Printf("char count of r is %d\n", utf8.RuneCountInString(s))
	}
	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}
}
