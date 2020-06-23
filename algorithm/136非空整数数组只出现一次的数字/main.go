package main

import (
	"github.com/google/go-cmp/cmp"
)

func main() {
	nums := []int{1, 2, 2, 1, 4, 5, 5}

	if !cmp.Equal(singleNumber(nums), 4) {
		panic("singleNumber not equal")
	}

	if !cmp.Equal(singleNumber2(nums), 4) {
		panic("singleNumber2 not equal")
	}
}

// 暴力法,hash实现,使用额外空间
func singleNumber(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		if _, exist := m[v]; exist {
			delete(m, v)
		} else {
			m[v] = 1
		}
	}
	for k := range m {
		return k
	}
	return 0
}

// 异域法,牛逼,没想到
func singleNumber2(nums []int) int {
	res := 0
	// for _, v := range nums {
	// 	res ^= v
	// }

	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}
