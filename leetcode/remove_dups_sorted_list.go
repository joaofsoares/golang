package leetcode

func deleteDuplicates(head *ListNode) *ListNode {

	res := &ListNode{}

	for head != nil {
		if head.Next != nil && head.Val != head.Next.Val {
			nodeAppend(res, head.Val)
		} else if head.Next == nil {
			nodeAppend(res, head.Val)
		}
		head = head.Next
	}

	return res.Next
}
