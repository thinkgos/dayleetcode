package main

import (
	"log"
)

// NOTE: 原地快排

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
	pivot := arr[0]
	for left < right {
		// 先找最右边的,大于pivot不动,直到找到小于pivot
		for left < right && arr[right] > pivot {
			right--
		}
		// 再找最左边的,小于pivot不动,直到找到大于pivot
		for left < right && arr[left] < pivot {
			left++
		}
		// 如果没碰到一起,就交换,
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	// 交换标轴点
	arr[0], arr[left] = arr[left], arr[0]

	// 分开排
	quickSort(arr[:left])
	quickSort(arr[left+1:])
}

func quickSort2(arr []int, left, right int) {
	if left >= right || len(arr) < 2 {
		return
	}

	l, r := left+1, right
	pivot := arr[left]
	for l < r {
		for l < r && arr[r] > pivot {
			r--
		}
		for l < r && arr[l] < pivot {
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
