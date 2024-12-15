package algorithm

// 1. 暴力解法
// 2. hash
func twoSum(nums []int, target int) []int {

	if len(nums) == 0 {
		return []int{}
	}

	mp := make(map[int]int, 0)
	for i, num := range nums {
		val, ok := mp[target-num]
		if ok {
			return []int{val, i}
		}
		mp[num] = i
	}

	return []int{}
}