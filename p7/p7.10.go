package main

import (
	"fmt"
)

func filtSlice(slice []int, fun func(int) bool) (newSlice []int) {
	newSlice = make([]int, 0)
	for _, val := range slice {
		if fun(val) {
			newSlice = append(newSlice, val)
		}
	}
	return
}

func main() {
	s := []int{1, 2, 3, 4, 6}
	fmt.Println(filtSlice(s, func(val int) bool {
		return 0 == val%2
	}))
}
