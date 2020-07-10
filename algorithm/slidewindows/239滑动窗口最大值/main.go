package main

import (
	"fmt"
)

func main() {
	input := []int{1, 4, 3, 2, 1, -3, 5, 3, 6}
	k := 3
	// output: [3,3,5,5,6,7]
	fmt.Printf("%#v\r\n", maxSlidingWindow(input, k))
	fmt.Printf("%#v\r\n", maxSlidingWindow1(input, k))
}

// 暴力法,轮一遍,在k中找出最大值
// 一共有L-K+1个窗口
func maxSlidingWindow(nums []int, k int) []int {
	length := len(nums)
	if length == 0 || k == 0 {
		return []int{}
	}
	win := make([]int, length-k+1)
	for i := 0; i < length-k+1; i++ {
		max := nums[i]
		for j := i + 1; j < i+k; j++ {
			if max < nums[j] {
				max = nums[j]
			}
		}
		win[i] = max
	}
	return win
}

// 双端队列法
func maxSlidingWindow1(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}
	queue := []int{} // 模拟一个双端队列
	result := make([]int, 0, len(nums)-k+1)
	for i := range nums {
		for i > 0 && (len(queue) > 0 && nums[i] > queue[len(queue)-1]) {
			// 将队列里比当前元素小的元素祭天,
			// 说白了,在它前面的都没它大,全砍死,在它后面的保留一波,排个小序
			queue = queue[:len(queue)-1]
		}
		// 将当前元素放入queue中
		queue = append(queue, nums[i])
		// 砍掉出窗口的最大元素
		// 维护队列,保证其头元素为当前窗口最大值
		if i >= k && nums[i-k] == queue[0] {
			queue = queue[1:]
		}

		// 已到窗口,开始保存最大元素,放入结果数组
		if i >= k-1 {
			result = append(result, queue[0])
		}
	}
	return result
}
