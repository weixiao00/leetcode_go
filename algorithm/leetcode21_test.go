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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	// -9,3
	// 5,7
	dummy := &ListNode{0, nil}
	tmp := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tmp.Next = list1
			list1 = list1.Next
		} else {
			tmp.Next = list2
			list2 = list2.Next
		}
		tmp = tmp.Next
	}

	for list1 != nil {
		tmp.Next = list1
		list1 = list1.Next
		tmp = tmp.Next
	}

	for list2 != nil {
		tmp.Next = list2
		list2 = list2.Next
		tmp = tmp.Next
	}

	return dummy.Next
}

func Test21(t *testing.T) {
	list1 := &ListNode{-9, &ListNode{3, nil}}
	list2 := &ListNode{5, &ListNode{7, nil}}
	res := mergeTwoLists(list1, list2)
	fmt.Println(res)
}
