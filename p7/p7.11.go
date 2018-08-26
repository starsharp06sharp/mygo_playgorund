package main

import (
	"fmt"
)

func insertStringSlice(dst []string, pos int, src []string) []string {
	return append(dst[:pos], append(src, dst[pos:]...)...)
}

func removeStringSlice(slice []string, start, end int) []string {
	return append(slice[:start], slice[end:]...)
}

func main() {
	sliceA := []string{"A", "B", "C"}
	sliceB := []string{"1", "2", "3"}
	newSlice := insertStringSlice(sliceA, 2, sliceB)
	fmt.Println(newSlice)
	fmt.Println(removeStringSlice(newSlice, 2, 5))
}
