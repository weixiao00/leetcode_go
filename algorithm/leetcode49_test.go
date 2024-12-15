package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

func groupAnagrams(strs []string) [][]string {

	if len(strs) == 0 {
		return [][]string{}
	}

	var posSort func(str string) string
	posSort = func(str string) string {
		bytes := []byte(str)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		return string(bytes)
	}

	posSortStrsMap := make(map[string][]string)
	for _, str := range strs {
		strs := posSortStrsMap[posSort(str)]
		if len(strs) == 0 {
			strs = make([]string, 0)
		}
		strs = append(strs, str)
		posSortStrsMap[posSort(str)] = strs
	}

	res := make([][]string, 0)
	for _, val := range posSortStrsMap {
		res = append(res, val)
	}

	return res
}

func Test49(t *testing.T) {
	strs := []string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"}
	res := groupAnagrams(strs)
	fmt.Println(res)
}
