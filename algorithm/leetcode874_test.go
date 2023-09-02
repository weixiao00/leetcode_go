package algorithm

import (
	"fmt"
	"strconv"
	"testing"
)

//设映射后的数字为kx+y，我们想知道k为多少时不会发生碰撞。当碰撞时kx1+y1=kx2+y2，
//解出来k=(y2-y1)/(x1-x2)，由于x,y的取值范围为[-30000,30000]，
//那么k的取值范围为[-60000,60000]，这个范围内的值都可能碰撞，比如k=60000时，(1,-30000)和(0,30000)算出来都是30000。
func robotSim(commands []int, obstacles [][]int) int {
	if commands == nil || obstacles == nil {
		return 0
	}
	arr := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	px, py := 0, 0
	point := 1
	res := 0
	obstaclesMap := map[string]interface{}{}

	for _, m := range obstacles {
		obstaclesMap[strconv.Itoa(m[0])+","+strconv.Itoa(m[1])] = 0
	}
	for _, step := range commands {
		if step < 0 {
			if step == -2 {
				point -= 1
			} else if step == -1 {
				point += 1
			}
			point %= 4
			if point < 0 {
				point += 4
			}
		} else {
			for i := 0; i < step; i++ {
				_, ok := obstaclesMap[strconv.Itoa(px+arr[point][0])+","+strconv.Itoa(py+arr[point][1])]
				if ok {
					break
				}
				px += arr[point][0]
				py += arr[point][1]
				if px*px+py*py > res {
					res = px*px + py*py
				}
			}
		}
	}
	return res
}

func Test874(t *testing.T) {
	commands := []int{4, -1, 4, -2, 4}
	obstacles := [][]int{{2, 4}}
	sim := robotSim(commands, obstacles)
	fmt.Println(sim)
}
