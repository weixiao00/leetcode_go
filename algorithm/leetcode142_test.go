package algorithm

// 解法一：快慢指针+hash存储
// 解法二：快慢指针+另一个节点从head走，与快慢指针相交点一起走相遇
// 解法三：hash存储，如果hash已经存在，说明是环入头
func detectCycle(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	slow := head
	fast := head
	var point *ListNode
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			point = slow
			break
		}
	}

	// 没有环
	if point == nil {
		return nil
	}

	nodeMap := make(map[*ListNode]bool)
	tmp := point.Next
	nodeMap[point] = true
	for tmp != point {
		nodeMap[tmp] = true
		tmp = tmp.Next
	}

	for head != nil {
		if nodeMap[head] {
			return head
		}
		head = head.Next
	}

	return nil
}

// hash存储
func detectCycle1(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	nodeMap := make(map[*ListNode]bool)

	for head != nil {
		if nodeMap[head] {
			return head
		}
		nodeMap[head] = true
		head = head.Next
	}
	return nil
}
