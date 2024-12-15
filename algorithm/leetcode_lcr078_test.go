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
//解法一：直接合并时间复杂度O(k^2*n)
//解法二：分治两两合并时间复杂度O(kn+lgk)
func mergeKLists(lists []*ListNode) *ListNode {

	if len(lists) == 0 {
		return nil
	}

	var merge func(l1, l2 *ListNode) *ListNode
	merge = func(l1, l2 *ListNode) *ListNode {
		if l1 == nil {
			return l2
		}
		if l2 == nil {
			return l1
		}
		dummy := &ListNode{0, nil}
		head := dummy
		for l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				dummy.Next = l1
				l1 = l1.Next
			} else {
				dummy.Next = l2
				l2 = l2.Next
			}
			dummy = dummy.Next
		}
		if l1 == nil {
			dummy.Next = l2
		} else {
			dummy.Next = l1
		}
		return head.Next
	}

	ans := lists
	for len(ans) > 1 {
		i := 0
		j := len(ans) - 1
		tmpAns := make([]*ListNode, 0)
		for i <= j {
			if i == j {
				tmpAns = append(tmpAns, ans[i])
				break
			}
			list := merge(ans[i], ans[j])
			i++
			j--
			tmpAns = append(tmpAns, list)
		}
		ans = tmpAns
	}
	return ans[0]
}

func TestLcr078(t *testing.T) {
	//[[-9,-7,-7],[-6,-4,-1,1],[-6,-5,-2,0,0,1,2],[-9,-8,-6,-5,-4,1,2,4],[-10],[-5,2,3]]
	l1_1 := &ListNode{-9, &ListNode{-7, &ListNode{-7, nil}}}
	l1_2 := &ListNode{-6, &ListNode{-4, &ListNode{-1, &ListNode{1, nil}}}}
	l1_3 := &ListNode{-6, &ListNode{-5, &ListNode{-2, &ListNode{0, &ListNode{0, &ListNode{1, &ListNode{2, nil}}}}}}}
	l1_4 := &ListNode{-9, &ListNode{-8, &ListNode{-6, &ListNode{-5, &ListNode{-4, &ListNode{1, &ListNode{2, &ListNode{4, nil}}}}}}}}
	l1_5 := &ListNode{-10, nil}
	l1_6 := &ListNode{-5, &ListNode{2, &ListNode{3, nil}}}

	lists := make([]*ListNode, 0)
	lists = append(lists, l1_1, l1_2, l1_3, l1_4, l1_5, l1_6)
	list := mergeKLists(lists)
	for list != nil {
		fmt.Print(list.Val)
		fmt.Print(" ")
		list = list.Next
	}
}
