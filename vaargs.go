package main

import "fmt"

func main() {
	printEveryln(1, "abc", 'a', 1+2i)
}

func printEveryln(argv ...interface{}) {
	for _, value := range argv {
		fmt.Println(value)
	}
}
