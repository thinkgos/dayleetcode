package main

import (
	"log"
)

func main() {
	result := []int{14, 9, 4, 10, 21, 12, 49, 8, 4, 3, 12, 20, 1, 200, 30, 50, 33, 22, 100, 45}
	selectSort(result)
	log.Printf("%#v\r\n", result)
}

// 选择排序,找出最小的,排前面.
func selectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		idx, min := i, arr[i]
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < min {
				min, idx = arr[j], j
			}
		}
		if idx != i {
			arr[i], arr[idx] = arr[idx], arr[i]
		}
	}
}
