package algorithm

import (
	"fmt"
	"testing"
)

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{0, head}

	pre := dummy
	for head != nil && head.Next != nil {
		next := head.Next.Next
		head.Next.Next = nil
		pre.Next = head.Next
		head.Next = next
		pre.Next.Next = head
		pre = head
		head = head.Next
	}

	return dummy.Next
}

func Test24(t *testing.T) {
	node4 := ListNode{4, nil}
	node3 := ListNode{3, &node4}
	node2 := ListNode{2, &node3}
	node1 := ListNode{1, &node2}
	pairs := swapPairs(&node1)
	fmt.Println(pairs)
}
