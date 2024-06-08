package algorithm

import (
	"container/list"
	"fmt"
	"testing"
)

func orangesRotting(grid [][]int) int {

	if len(grid) == 0 {
		return 0
	}

	queue := list.New()

	containOne := false
	for i, v := range grid {
		for j, vv := range v {
			if vv == 2 {
				queue.PushBack([]int{i, j})
			}
			if vv == 1 {
				containOne = true
			}
		}
	}

	if queue.Len() == 0 {
		if containOne {
			return -1
		}
		return 0
	}

	time := -1
	dr := []int{1, 0, -1, 0}
	dc := []int{0, -1, 0, 1}
	totalRow := len(grid)
	totalColumn := len(grid[0])
	for queue.Len() != 0 {

		len := queue.Len()
		time++
		for k := 0; k < len; k++ {
			// 取出来队列头节点
			orange := queue.Remove(queue.Front()).([]int)
			for i := 0; i < 4; i++ {
				row := orange[0]
				column := orange[1]
				row += dr[i]
				column += dc[i]
				if row >= 0 && row < totalRow && column >= 0 && column < totalColumn && grid[row][column] == 1 {
					queue.PushBack([]int{row, column})
					grid[row][column] = 2
				}
			}
		}
	}

	for _, v := range grid {
		for _, vv := range v {
			if vv == 1 {
				time = -1
			}
		}
	}

	return time
}

func Test994(t *testing.T) {
	grid := [][]int{{0}}
	rotting := orangesRotting(grid)
	fmt.Println(rotting)
}
