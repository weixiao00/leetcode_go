package algorithm

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

// 循环获取下一个排列
func permuteUnique(nums []int) [][]int {

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
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] < nums[i+1] {
				p = i
			}
		}
		if p == -1 {
			return res
		}
		q := -1
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

// dfs递归。通过map的key从数组变成字符串去重
func permuteUnique1(nums []int) [][]int {

	if len(nums) == 0 {
		return [][]int{}
	}

	mp := make(map[string][]int, 0)
	visited := make([]bool, len(nums))

	var intConvStr func(arr []int) string
	intConvStr = func(arr []int) string {
		res := ""
		for _, val := range arr {
			res += strconv.Itoa(val)
		}
		return res
	}

	var dfs func(curRes []int, idx int)
	dfs = func(curRes []int, idx int) {
		if idx == len(nums) {
			copyRes := make([]int, len(curRes))
			copy(copyRes, curRes)
			mp[intConvStr(curRes)] = copyRes
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
	res := make([][]int, 0)
	for _, val := range mp {
		res = append(res, val)
	}
	return res
}

// 官方题解
func permuteUnique2(nums []int) [][]int {

	if len(nums) == 0 {
		return [][]int{}
	}

	// 先排序
	sort.Ints(nums)
	res := make([][]int, 0)
	visited := make([]bool, len(nums))
	var dfs func(idx int, curRes []int)
	dfs = func(idx int, curRes []int) {
		if idx == len(nums) {
			// 创建一个新数组
			res = append(res, append([]int{}, curRes...))
			return
		}

		for i := range nums {
			if visited[i] || (i > 0 && nums[i] == nums[i-1] && !visited[i-1]) {
				continue
			}
			visited[i] = true
			dfs(idx+1, append(curRes, nums[i]))
			visited[i] = false
		}
	}

	dfs(0, []int{})
	return res
}

func Test47(t *testing.T) {
	nums := []int{1, 1, 2}
	res := permuteUnique2(nums)
	fmt.Println(res)
	// 使用append切片不一定会产生一个新的数组。看容量是否充足，如果充足的话，会在原来的数组上追加
	// 这样的话虽然引用的是同一个数组。但是len是不一样的
	//arr := make([]int, 3, 4)
	//arr1 := append(arr, 5)
	////arr1[0] = 100
	//arr2 := append(arr, 6)
	//fmt.Println(arr)
	//fmt.Println(arr1)
	//fmt.Println(arr2)
}
