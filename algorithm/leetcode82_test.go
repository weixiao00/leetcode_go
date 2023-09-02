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

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{0, head}
	tmp := dummy
	for dummy.Next != nil && dummy.Next.Next != nil {
		val := dummy.Next.Val
		if val == dummy.Next.Next.Val {
			for dummy.Next != nil && val == dummy.Next.Val {
				dummy.Next = dummy.Next.Next
			}
		} else {
			dummy = dummy.Next
		}
	}
	return tmp.Next
}

func TestDeleteDuplicates(t *testing.T) {
	node3 := ListNode{3, nil}
	node2 := ListNode{1, &node3}
	node1 := ListNode{1, &node2}

	duplicates := deleteDuplicates(&node1)
	for duplicates != nil {
		fmt.Println(duplicates.Val)
		duplicates = duplicates.Next
	}
}
