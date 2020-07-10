package main

import (
	"fmt"
)

func main() {
	// 27
	var nums = []int{3, 2, 2, 3}
	var val = 3
	length := removeElement(nums, val)
	fmt.Printf("length: %d\r\n", length)

	nums = []int{3, 2, 2, 3, 4, 2, 1, 5, 4}
	length = removeElement1(nums, val)
	fmt.Printf("length: %d\r\n", length)

	// 26
	var numsSort = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	length = removeDuplicates(numsSort)
	fmt.Printf("length: %d\r\n", length)
}

// 27 题 方法1
func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}

// 27 题 方法2
func removeElement1(nums []int, val int) int {
	for i := 0; i < len(nums); {
		if nums[i] == val {
			nums[i] = nums[len(nums)-1]
			nums = nums[:len(nums)-1]
		} else {
			i++
		}
	}
	fmt.Printf("%#v --> %d\r\n", nums, len(nums))
	return len(nums)
}

// 26题 在排序数组,原地删除重复出现的元素,使每个元素只出现一次
func removeDuplicates(nums []int) int {
	for i := 0; i+1 < len(nums); {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}
