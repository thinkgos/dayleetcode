package main

import (
	"testing"
)

func BenchmarkRemoveElement(b *testing.B) {
	var nums = []int{3, 2, 2, 3}
	var val = 3

	for i := 0; i < b.N; i++ {
		removeElement(nums, val)
	}
}

func BenchmarkRemoveElement1(b *testing.B) {
	var nums = []int{3, 2, 2, 3}
	var val = 3

	for i := 0; i < b.N; i++ {
		removeElement1(nums, val)
	}
}
