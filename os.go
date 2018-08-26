package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	var goos string = runtime.GOOS
	fmt.Printf("This os is %s\n", goos)
	envPath := os.Getenv("PATH")
	fmt.Printf("Current path is %s\n", envPath)
}
