package main

import (
	"fmt"
)

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6}
	var k = 3
	// 1. 暴力法
	//  直接每个元素右移k位,不支持原地移动

	// 2. 三次原地反转法
	val := rotate(nums, k)
	fmt.Printf("%+v", val)
}

func rotate(nums []int, k int) []int {
	reverse(nums)
	reverse(nums[:k%len(nums)])
	reverse(nums[k%len(nums):])
	return nums
}

func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}
