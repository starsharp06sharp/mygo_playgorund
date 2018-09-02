package main

import "fmt"

func main() {
	var someMap = map[int]string{1: "a", 2: "b"}

	fmt.Println(someMap)
	fmt.Printf("%v\n", someMap)
	// %v	the value in a default format
	// when printing structs, the plus flag (%+v) adds field names
	// %#v	a Go-syntax representation of the value
	// %T	a Go-syntax representation of the type of the value
	// %%	a literal percent sign; consumes no value

	val, ok := someMap[1]
	fmt.Println(val, ok)

	val, ok = someMap[3]
	fmt.Println(val, ok)

	val, ok = someMap[3]
	fmt.Println(val, ok)
}
