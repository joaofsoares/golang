package leetcode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	tmpNode := head

	cnt := 1
	for tmpNode.Next != nil {
		tmpNode = tmpNode.Next
		cnt++
	}
	tmpNode.Next = head

	length := cnt - k%cnt

	for i := 0; i < length; i++ {
		tmpNode = tmpNode.Next
	}

	head = tmpNode.Next
	tmpNode.Next = nil

	return head
}
