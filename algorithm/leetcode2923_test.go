package algorithm

func findChampion(grid [][]int) int {

	if len(grid) == 0 {
		return 0
	}
	for i, nums := range grid {
		count := 0
		for _, num := range nums {
			if num == 1 {
				count++
			}
		}
		if count == len(grid)-1 {
			return i
		}
	}

	return 0
}