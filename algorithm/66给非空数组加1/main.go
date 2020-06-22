package main

func plusOne(digits []int) []int {
	for i := len(digits); i >= 0; i-- {
		if digits[i] < 9 { // 无需进位
			digits[i]++
			return digits
		} else { // 需要进位
			digits[i] = 0
		}
	}
	return append([]int{1}, digits...)
}
