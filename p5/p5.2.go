package main

import "fmt"

// Season 函数返回月份对应季节的中文名字
func Season(month int) (string, bool) {
	switch month {
	case 2, 3, 4:
		return "春", true
	case 5, 6, 7:
		return "夏", true
	case 8, 9, 10:
		return "秋", true
	case 11, 12, 1:
		return "冬", true
	default:
		return "", false
	}
}

func main() {
	var month int
	fmt.Scanf("%d", &month)
	res, ok := Season(month)
	if ok {
		fmt.Printf("%d month is %s\n", month, res)
	} else {
		fmt.Printf("Error! Invalid month:%d\n", month)
	}
}
