package algorithm

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	dummy := &ListNode{}
	tail := dummy
	// 进位
	carry := 0
	for l1 != nil || l2 != nil {
		v1 := 0
		if l1 != nil {
			v1 = l1.Val
		}
		v2 := 0
		if l2 != nil {
			v2 = l2.Val
		}
		curVal := v1 + v2 + carry
		tail.Next = &ListNode{curVal % 10, nil}
		tail = tail.Next
		carry = curVal / 10

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry == 1 {
		tail.Next = &ListNode{1, nil}
	}

	return dummy.Next
}

