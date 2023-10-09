package algorithm

import (
	"fmt"
	"testing"
)

// 拓扑排序
func canFinish(numCourses int, prerequisites [][]int) bool {

	if len(prerequisites) == 0 {
		return true
	}

	ingree := make([]int, numCourses)
	// 计算入度
	for _, arr := range prerequisites {
		ingree[arr[0]] += 1
	}

	queue := make([]int, 0)
	// 入度为0
	for i, _ := range ingree {
		if ingree[i] == 0 {
			queue = append(queue, i)
		}
	}

	if len(queue) == 0 {
		return false
	}

	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		for _, arr := range prerequisites {
			if arr[1] == i {
				ingree[arr[0]] -= 1
				if ingree[arr[0]] == 0 {
					queue = append(queue, arr[0])
				}
			}
		}
	}

	// 还有节点度大于0
	for i := range ingree {
		if ingree[i] > 0 {
			return false
		}
	}

	return true
}

func Test207(t *testing.T) {
	numCourses := 4
	prerequisites := [][]int{{1, 0}, {2, 1}, {3, 2}}
	finish := canFinish(numCourses, prerequisites)
	fmt.Println(finish)
}
