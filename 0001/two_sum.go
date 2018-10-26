package twoSum

// 8ms 54.7% chris
func twoSum(nums []int, target int) []int {
	// 最快时间4ms，区别在于make
	// m := map[int]int{} -- 4ms
	m := make(map[int]int)

	for i, v := range nums {
		check := target - v
		if _, ok := m[check]; ok {
			if m[check] != i {
				return []int{i, m[check]}
			}
		} else {
			m[v] = i
		}
	}

	return []int{}

}
