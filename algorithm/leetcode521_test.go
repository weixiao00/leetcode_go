package algorithm

func findLUSlength(a string, b string) int {

	if a == b {
		return -1
	}
	return getMax13(len(a), len(b))
}

func getMax13(a, b int) int {
	if a > b {
		return a
	}
	return b
}
