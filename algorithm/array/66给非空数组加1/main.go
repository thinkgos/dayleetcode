package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%#v\r\n", plusOne([]int{1, 9, 8}))
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 { // 无需进位
			digits[i]++
			return digits
		} else { // 需要进位
			digits[i] = 0
		}
	}
	return append([]int{1}, digits...)
}
