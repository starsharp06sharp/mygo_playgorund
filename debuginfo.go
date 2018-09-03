package main

import (
	"fmt"
	"log"
	"runtime"
)

func printCaller() {
	fmt.Println(runtime.Caller(1))
}

func main() {
	printCaller()
	log.SetFlags(log.Llongfile)
	log.Print("")
}
