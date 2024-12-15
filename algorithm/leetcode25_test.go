package algorithm

import (
	"fmt"
	"testing"
)

// k个一组链表反转
func reverseKGroup(head *ListNode, k int) *ListNode {

	if head == nil || k <= 1 {
		return head
	}

	var reverse func(node *ListNode) *ListNode
	reverse = func(node *ListNode) *ListNode {
		if node == nil {
			return node
		}

		var pre *ListNode = nil
		for node != nil {
			next := node.Next
			node.Next = pre
			pre = node
			node = next
		}
		return pre
	}

	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	for pre != nil {

		next := pre
		for i := 0; i < k; i++ {
			next = next.Next
			if next == nil {
				return dummy.Next
			}
		}

		// 将k个链表和后面的断开
		p := next.Next
		next.Next = nil
		next = p

		newHead := reverse(pre.Next)
		tail := pre.Next
		tail.Next = next
		pre.Next = newHead
		pre = tail
	}

	return dummy.Next
}

func Test25(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	dummy := &ListNode{}
	tmp := dummy
	for _, nums := range nums {
		tmp.Next = &ListNode{nums, nil}
		tmp = tmp.Next
	}
	k := 3
	res := reverseKGroup(dummy.Next, k)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
