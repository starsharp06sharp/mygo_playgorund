package main

import (
	"fmt"
	"sort"
)

func main() {
	var en2cnDrink = map[string]string{
		"juice":  "果汁",
		"coke":   "可乐",
		"fanta":  "芬达",
		"sprite": "雪碧",
	}
	fmt.Print("Drink name: ")
	for enName := range en2cnDrink {
		fmt.Print(enName, ", ")
	}
	fmt.Print("\n")

	fmt.Print("Drink name 2cn: ")
	for enName, cnName := range en2cnDrink {
		fmt.Print(enName, ":", cnName, ", ")
	}
	fmt.Print("\n")

	var enNameList = make([]string, len(en2cnDrink))
	var i = 0
	for enName := range en2cnDrink {
		enNameList[i] = enName
		i++
	}
	sort.Strings(enNameList)
	fmt.Print("Sorted drink name 2cn: ")
	for _, enName := range enNameList {
		fmt.Print(enName, ":", en2cnDrink[enName], ", ")
	}
	fmt.Print("\n")
}
