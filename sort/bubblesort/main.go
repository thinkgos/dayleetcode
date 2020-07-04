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
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
