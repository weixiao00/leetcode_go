package algorithm

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	if len(hours) == 0 {
		return 0
	}
	res := 0
	for _, h := range hours {
		if h >= target {
			res++
		}
	}

	return res
}
