package algorithm

func imageSmoother(img [][]int) [][]int {

	if len(img) == 0 {
		return [][]int{}
	}

	m := len(img)
	n := len(img[0])

	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			num := 0
			sum := 0
			for x := i - 1; x <= i+1; x++ {
				for y := j - 1; y <= j+1; y++ {
					if x >= 0 && x < m && y >= 0 && y < n {
						num++
						sum += img[x][y]
					}
				}
			}

			res[i][j] = sum / num
		}
	}

	return res
}
