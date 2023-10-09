package algorithm

// 拓扑排序
func findOrder(numCourses int, prerequisites [][]int) []int {

	init := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		init = append(init, i)
	}

	if len(prerequisites) == 0 {
		return init
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
		return []int{}
	}

	res := make([]int, 0)
	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		res = append(res, i)
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
			return []int{}
		}
	}

	return res
}
