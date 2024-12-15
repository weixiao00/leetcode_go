package algorithm

import "testing"

// 方法一，使用栈
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}

	stack1 := make([]*ListNode, 0)
	stack2 := make([]*ListNode, 0)

	for headA != nil {
		stack1 = append(stack1, headA)
		headA = headA.Next
	}
	for headB != nil {
		stack2 = append(stack2, headB)
		headB = headB.Next
	}

	var node *ListNode = nil
	for len(stack1) != 0 && len(stack2) != 0 {
		node1 := stack1[len(stack1)-1]
		stack1 = stack1[0 : len(stack1)-1]
		node2 := stack2[len(stack2)-1]
		stack2 = stack2[0 : len(stack2)-1]
		if node1 != node2 {
			return node
		}
		node = node1
	}

	return node
}

// 方法二，快慢指针
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	tmp1 := headA
	count1 := 0
	for tmp1 != nil {
		count1++
		tmp1 = tmp1.Next
	}

	tmp2 := headB
	count2 := 0
	for tmp2 != nil {
		count2++
		tmp2 = tmp2.Next
	}
	if count1 > count2 {
		for i := 0; i < count1-count2; i++ {
			headA = headA.Next
		}
	} else {
		for i := 0; i < count2-count1; i++ {
			headB = headB.Next
		}
	}

	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}

	return nil
}

// 方法三，太暴力
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	tmp1 := headA
	tmp2 := headB
	for tmp1 != nil {
		for tmp2 != nil {
			if tmp1 == tmp2 {
				return tmp1
			}
			tmp2 = tmp2.Next
		}
		tmp2 = headB
		tmp1 = tmp1.Next
	}

	return nil
}

func Test160(t *testing.T) {

}
