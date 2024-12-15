package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

// 解法一：map然后排序  O(nlgn)
// 解法二：map统计数量，k个大小的小顶堆。跟堆顶比较 (nlgk)
// 解法三：桶排序。相同的出现频率放入一个桶中 O(n)
func topKFrequent(nums []int, k int) []int {

	if len(nums) == 0 {
		return []int{}
	}

	mp := make(map[int]int)
	for i := range nums {
		mp[nums[i]] = mp[nums[i]] + 1
	}

	type kv struct {
		key int
		val int
	}

	kvs := make([]kv, 0)
	for key, val := range mp {
		kvs = append(kvs, kv{key: key, val: val})
	}
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].val > kvs[j].val
	})

	res := make([]int, 0)
	for i := 0; i < k; i++ {
		res = append(res, kvs[i].key)
	}

	return res
}

// 桶排序
func topKFrequent1(nums []int, k int) []int {

	if len(nums) == 0 {
		return []int{}
	}

	mp := make(map[int]int)
	for i := range nums {
		mp[nums[i]] = mp[nums[i]] + 1
	}

	list := make([][]int, len(nums)+1)
	for key, val := range mp {
		list[val] = append(list[val], key)
	}

	res := make([]int, 0)
	for i := len(list) - 1; i >= 0 && len(res) < k; i-- {
		if len(list[i]) == 0 {
			continue
		}
		res = append(res, list[i]...)
	}

	return res
}

func Test347(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	res := topKFrequent1(nums, k)
	fmt.Println(res)
}
