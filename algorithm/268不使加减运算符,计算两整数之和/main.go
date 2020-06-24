package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(100, 32))
}

func sum(a, b int) int {
	for b != 0 {
		temp := a ^ b
		b = (a & b) << 1
		a = temp
	}
	return a
}
