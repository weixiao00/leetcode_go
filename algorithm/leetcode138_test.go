package algorithm

import (
	"fmt"
	"testing"
)

/**
* Definition for a Node.
* type Node struct {
*     Val int
*     Next *Node
*     Random *Node
* }
 */

type Node1 struct {
	Val    int
	Next   *Node1
	Random *Node1
}

// 复制出来多个next节点
func copyRandomList(head *Node1) *Node1 {

	if head == nil {
		return head
	}

	// 复制节点
	tmp := head
	for tmp != nil {
		next := &Node1{tmp.Val, tmp.Next, nil}
		tmp.Next = next
		tmp = tmp.Next.Next
	}

	// 连接random
	tmp = head
	for tmp != nil {
		if tmp.Random != nil {
			next := tmp.Next
			next.Random = tmp.Random.Next
		}
		tmp = tmp.Next.Next
	}

	// 构造新链表.还原原链表
	tmp = head
	dummy := &Node1{-1, nil, nil}
	pre := dummy
	for tmp != nil {
		pre.Next = tmp.Next
		pre = pre.Next
		// 还原原链表
		tmp.Next = pre.Next
		tmp = pre.Next
	}

	return dummy.Next
}

func Test138(t *testing.T) {
	head := &Node1{1, nil, nil}
	head2 := &Node1{2, nil, nil}
	head3 := &Node1{3, nil, nil}
	head.Next = head2
	head2.Next = head3
	head.Random = head3
	head3.Random = head
	res := copyRandomList(head)
	fmt.Println(res)
}
