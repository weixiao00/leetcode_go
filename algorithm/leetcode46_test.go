package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

// 获取下一个排列，循环
func permute(nums []int) [][]int {

	if len(nums) == 0 {
		return [][]int{}
	}

	res := make([][]int, 0)
	sort.Ints(nums)
	copyNums := make([]int, len(nums))
	copy(copyNums, nums)
	res = append(res, copyNums)

	for true {

		p := -1
		// 寻找往右边的最小数
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] < nums[i+1] {
				p = i
			}
		}
		if p == -1 {
			return res
		}
		q := -1
		// 寻找大于p的右边最小数
		for j := p + 1; j < len(nums); j++ {
			if nums[p] < nums[j] {
				q = j
			}
		}

		// 交换
		tmp := nums[p]
		nums[p] = nums[q]
		nums[q] = tmp

		// 交换排序
		i := p + 1
		j := len(nums) - 1
		for i < j {
			tmp := nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
			i++
			j--
		}
		copyNums := make([]int, len(nums))
		copy(copyNums, nums)
		res = append(res, copyNums)
	}

	return res
}

// dfs递归
func permute1(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	res := make([][]int, 0)
	visited := make([]bool, len(nums))

	var dfs func(curRes []int, idx int)
	dfs = func(curRes []int, idx int) {
		if idx == len(nums) {
			// 这里最好是copy一个数组，因为这里传的引用，可能是递归上一个的相同的数组。看切片的容量
			res = append(res, curRes)
			return
		}

		for i := range nums {
			if visited[i] {
				continue
			}
			visited[i] = true
			dfs(append(curRes, nums[i]), idx+1)
			visited[i] = false
		}
	}

	dfs([]int{}, 0)

	return res
}

func Test46(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	res := permute1(nums)
	fmt.Println(res)
}
