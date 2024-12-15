package algorithm

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//解法一：直接合并时间复杂度O(k^2*n)
//解法二：分治两两合并时间复杂度O(kn+lgk)。归并排序合并
//与lcr078相同的题目
func mergeKLists1(lists []*ListNode) *ListNode {

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
