package main

import (
	"fmt"
	"math/big"
)

func jc(num int) *big.Int {
	if 2 >= num {
		return big.NewInt(int64(num))
	}
	res := big.NewInt(int64(num))
	return res.Mul(res, jc(num-1))
}

func main() {
	var num int
	fmt.Scanf("%d", &num)
	for i := 1; i <= num; i++ {
		fmt.Printf("%d!=%d\n", i, jc(i))
	}
}
