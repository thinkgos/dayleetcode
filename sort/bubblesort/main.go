package main

import (
	"log"
)

func main() {
	result := []int{14, 9, 4, 10, 21, 12, 49, 8, 4, 3, 12, 20, 1, 200, 30, 50, 33, 22, 100, 45}
	bubbleSort(result)
	log.Printf("%#v\r\n", result)
}

func bubbleSort(arr []int) {
	// go 中,其实这两可以不判断
	if arr == nil || len(arr) < 2 {
		return
	}

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
