package main

import (
	"fmt"
)

func mapFun(fun func(int) int, lst []int) (res []int) {
	res = make([]int, len(lst))
	for i, val := range lst {
		res[i] = fun(val)
	}
	return
}

func main() {
	lst := []int{1, 3, 5, 7, 9}
	fmt.Println(mapFun(func(val int) int {
		return val * 10
	}, lst))
}
