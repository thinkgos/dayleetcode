package main

func intersect(num1, num2 []int) []int {
	m := make(map[int]int)

	for _, v := range num1 {
		m[v] += 1
	}
	result := num1[:0]
	for _, v := range num2 {
		if m[v] > 0 {
			m[v] -= 1
			result = append(result, v)
		}
	}
	return result
}
