package main

import (
	"log"
)

func main() {
	result := []int{14, 9, 4, 10, 21, 12, 49, 8, 4, 3, 12, 20}
	//quickSort(result)
	quickSort2(result, 0, len(result)-1)
	log.Printf("%#v\r\n", result)
}

func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	left, right := 1, len(arr)-1
	privot := arr[0]
	for left < right {
		for left < right && arr[right] > privot {
			right--
		}
		for left < right && arr[left] < privot {
			left++
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	arr[0], arr[left] = arr[left], arr[0]

	quickSort(arr[:left])
	quickSort(arr[left+1:])
}

func quickSort2(arr []int, left, right int) {
	if left >= right || len(arr) < 2 {
		return
	}

	l, r := left+1, right
	privot := arr[left]
	for l < r {
		for l < r && arr[r] > privot {
			r--
		}
		for l < r && arr[l] < privot {
			l++
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}
	arr[left], arr[l] = arr[l], arr[left]

	quickSort2(arr, left, l-1)
	quickSort2(arr, l+1, right)
}
