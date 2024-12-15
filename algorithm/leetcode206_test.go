package algorithm

import (
	"fmt"
	"testing"
)

func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return head
	}

	var pre *ListNode = nil
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}

	return pre
}

func Test206(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	dummy := &ListNode{}
	tmp := dummy
	for _, nums := range nums {
		tmp.Next = &ListNode{nums, nil}
		tmp = tmp.Next
	}
	res := reverseList(dummy.Next)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
