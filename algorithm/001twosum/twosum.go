package _01twosum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for idx2, v := range nums {
		if idx1, ok := m[target-v]; ok {
			return []int{idx1, idx2}
		}
		m[v] = idx2
	}
	return []int{}
}
