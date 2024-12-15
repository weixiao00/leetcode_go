package algorithm

func timeRequiredToBuy(tickets []int, k int) int {

	if len(tickets) == 0 {
		return 0
	}

	var getMin func(a, b int) int
	getMin = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	res := 0
	for i, ticket := range tickets {
		if i <= k {
			res += getMin(ticket, tickets[k])
		} else {
			res += getMin(ticket, tickets[k]-1)
		}
	}

	return res
}
