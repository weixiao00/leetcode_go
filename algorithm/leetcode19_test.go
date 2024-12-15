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
// 构造哑节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head == nil {
		return nil
	}

	// 获取链表的长度
	tmp := head
	l := 0
	for tmp != nil {
		tmp = tmp.Next
		l++
	}

	k := l - n
	dummy := &ListNode{0, head}
	tmp = dummy
	for i := 0; i < k; i++ {
		tmp = tmp.Next
	}
	tmp.Next = tmp.Next.Next
	return dummy.Next
}

func Test19(t *testing.T) {
	head := &ListNode{1, &ListNode{2, nil}}
	res := removeNthFromEnd(head, 2)
	fmt.Println(res)
}
