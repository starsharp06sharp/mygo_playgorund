package main

import (
	"fmt"
)

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		flag := true
		for j := 0; j < n-i-1; j++ {
			if arr[j] < arr[j+1] {
				flag = false
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		if flag {
			break
		}
	}
}

func main() {
	arr := []int{1, 3, 5, 2, 4}
	bubbleSort(arr)
	fmt.Println(arr)
}
