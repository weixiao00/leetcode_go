package algorithm

// 双指针
func trap(height []int) int {
	if height == nil {
		return 0
	}
	res := 0
	begin := 0
	end := len(height) - 1
	leftMax := height[begin]
	rightMax := height[end]
	for begin < end {
		if leftMax < rightMax {
			if leftMax > height[begin+1] {
				res += leftMax - height[begin+1]
			} else {
				leftMax = height[begin+1]
			}
			begin++
		} else {
			if rightMax > height[end-1] {
				res += rightMax - height[end-1]
			} else {
				rightMax = height[end-1]
			}
			end--
		}
	}
	return res
}
