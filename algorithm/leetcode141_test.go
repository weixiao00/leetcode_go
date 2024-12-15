package algorithm

import "testing"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 快慢指针
// 不知道有啥问题，但是lc过不了
func hasCycle(head *ListNode) bool {

	if head == nil {
		return false
	}

	fast := head
	slow := head

	// 使用慢指针作为判断条件是有问题的
	for slow != nil {
		slow = slow.Next
		if fast != nil && fast.Next != nil {
			fast = fast.Next.Next
		}
		if fast != nil && slow == fast {
			return true
		}
	}

	return false
}

func hasCycle1(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast := head
	slow := head
	// 判断快指针是否为空
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func Test141(t *testing.T) {

}
