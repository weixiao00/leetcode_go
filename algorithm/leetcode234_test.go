package algorithm

import (
	"fmt"
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 解法一：反转链表
// 解法二：数据存储，双指针
// 解法三：递归。空间复杂度最低
func isPalindrome1(head *ListNode) bool {

	sourceHead := head
	var recursivelyCheck func(head *ListNode) bool
	recursivelyCheck = func(head *ListNode) bool {
		if head == nil {
			return true
		}
		if !recursivelyCheck(head.Next) {
			return false
		}
		if head.Val != sourceHead.Val {
			return false
		}
		sourceHead = sourceHead.Next
		return true
	}

	return recursivelyCheck(head)
}

func Test234(t *testing.T) {

	head3 := &ListNode{1, nil}
	head2 := &ListNode{2, head3}
	head1 := &ListNode{2, head2}
	head := &ListNode{1, head1}
	res := isPalindrome1(head)
	fmt.Println(res)
}
