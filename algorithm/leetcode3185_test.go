package algorithm

func countCompleteDayPairs1(hours []int) int64 {

	if len(hours) == 0 {
		return 0
	}

	res := 0
	cnt := make([]int, 24)
	for _, hour := range hours {
		res += cnt[(24-hour%24)%24]
		cnt[hour%24]++
	}

	return int64(res)
}
