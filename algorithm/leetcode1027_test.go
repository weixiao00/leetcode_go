package algorithm

import "testing"

/**
给你一个整数数组nums，返回 nums中最长等差子序列的长度。

回想一下，nums 的子序列是一个列表 nums[i1], nums[i2], ..., nums[ik] ，且 0 <= i1 < i2 < ... < ik <= nums.length - 1。并且如果 seq[i+1] - seq[i]( 0 <= i < seq.length - 1) 的值都相同，那么序列 seq 是等差的。

*/
func Test1027(t *testing.T) {

}

func longestArithSeqLength(nums []int) int {
	return 0
}

//class Solution {
//    public int longestArithSeqLength(int[] nums) {
//        int minv = Arrays.stream(nums).min().getAsInt();
//        int maxv = Arrays.stream(nums).max().getAsInt();
//        int diff = maxv - minv;
//        int ans = 1;
//        for (int d = -diff; d <= diff; ++d) {
//            int[] f = new int[maxv + 1];
//            Arrays.fill(f, -1);
//            for (int num : nums) {
//                int prev = num - d;
//                if (prev >= minv && prev <= maxv && f[prev] != -1) {
//                    f[num] = Math.max(f[num], f[prev] + 1);
//                    ans = Math.max(ans, f[num]);
//                }
//                f[num] = Math.max(f[num], 1);
//            }
//        }
//        return ans;
//    }
//}
//
