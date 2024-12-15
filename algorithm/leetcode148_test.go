package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	valList := make([]int, 0)
	tmp := head
	for tmp != nil {
		valList = append(valList, tmp.Val)
		tmp = tmp.Next
	}
	sort.Ints(valList)
	dummy := &ListNode{-1, nil}
	tmp = dummy
	for _, val := range valList {
		tmp.Next = &ListNode{val, nil}
		tmp = tmp.Next
	}

	return dummy.Next
}

func Test148(t *testing.T) {
	//head = [4,2,1,3]
	head := &ListNode{4, nil}
	head1 := &ListNode{2, nil}
	head2 := &ListNode{1, nil}
	head3 := &ListNode{3, nil}
	head.Next = head1
	head1.Next = head2
	head2.Next = head3
	list := sortList(head)
	fmt.Println(list)
}
