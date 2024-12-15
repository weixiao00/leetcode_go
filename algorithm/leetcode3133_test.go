package algorithm

import (
	"fmt"
	"testing"
)

//public long minEnd(int n, int x) {
//        long ans = x;
//        int m = n - 1;
//        for (long i = 1, offset = 0; i <= m; i <<= 1) {
//            while ((ans & (i << offset)) > 0) {
//                offset++;
//            }
//            ans |= (m & i) << offset;
//        }
//        return ans;
//    }
// 没太理解呢？？？
func minEnd(n int, x int) int64 {

	ans := x
	m := n - 1

	offset := 0
	for i := 1; i <= m; i <<= 1 {
		for (ans & (i << offset)) > 0 {
			offset++
		}
		ans |= (m & i) << offset
	}

	return int64(ans)
}

func Test3133(t *testing.T) {
	n := 3
	x := 1
	res := minEnd(n, x)
	fmt.Println(res)
}
