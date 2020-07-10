package main

import (
	"fmt"
)

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(findMissNumber(nums))
	fmt.Println(findMissNumber1(nums))
}

// 高斯方程法
func findMissNumber(nums []int) int {
	result := len(nums) * (len(nums) + 1) / 2
	for _, v := range nums {
		result -= v
	}
	return result
}

// 位法
func findMissNumber1(nums []int) int {
	var result = 0

	for i, v := range nums {
		result ^= i ^ v
	}
	return result ^ len(nums)
}
