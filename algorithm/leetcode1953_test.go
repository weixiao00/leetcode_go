package algorithm

func numberOfWeeks(milestones []int) int64 {
	if len(milestones) == 0 {
		return 0
	}
	// 找到最大的周数
	maxM := milestones[0]
	rest := 0
	for _, milestone := range milestones {
		rest += milestone
		maxM = getMax2(maxM, milestone)
	}

	// 除去最大的所有其余的
	rest -= maxM
	// 其余的大于等于最大的，都可以完成
	if maxM <= rest+1 {
		return int64(rest + maxM)
	} else {
		// 完成一部分
		return int64(2*rest + 1)
	}
}

func getMax2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
