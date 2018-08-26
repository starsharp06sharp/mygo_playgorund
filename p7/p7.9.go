package main

import (
	"fmt"
)

func expandSlice(slice []int, factor int) []int {
	newSlice := make([]int, 0)
	for ; 0 != factor; factor /= 2 {
		if 1 == factor%2 {
			newSlice = append(newSlice, slice...)
		}
		slice = append(slice, slice...)
	}
	return newSlice
}

func main() {
	newSlice := expandSlice([]int{1, 2, 3}, 3)
	fmt.Println(len(newSlice))
}
