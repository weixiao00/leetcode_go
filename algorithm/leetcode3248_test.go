package algorithm

// 模拟
//"UP"、"RIGHT"、"DOWN" 和 "LEFT"
func finalPositionOfSnake(n int, commands []string) int {

	if len(commands) == 0 {
		return 0
	}

	i := 0
	j := 0
	for _, command := range commands {
		if command == "UP" {
			i--
		} else if command == "DOWN" {
			i++
		} else if command == "LEFT" {
			j--
		} else if command == "RIGHT" {
			j++
		}
	}

	return i*n + j
}
