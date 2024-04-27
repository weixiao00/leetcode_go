package algorithm

func maximumBinaryString(binary string) string {
	if len(binary) == 0 {
		return ""
	}
	j := 0
	s := []rune(binary)
	for i := 0; i < len(binary); i++ {
		if s[i] == '0' {
			for j <= i || (j < len(binary) && s[j] == '1') {
				j++
			}
			if j < len(binary) {
				s[i] = '1'
				s[j] = '1'
				s[i+1] = '0'
			}
		}
	}
	return string(s)
}
