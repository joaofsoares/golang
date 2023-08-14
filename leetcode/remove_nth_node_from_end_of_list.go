package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodeAppend(n *ListNode, val int) {
	end := &ListNode{Val: val}
	here := n
	for here.Next != nil {
		here = here.Next
	}
	here.Next = end
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	node := head

	var newNode ListNode

	cnt := 0

	for node != nil {
		nodeAppend(&newNode, node.Val)
		node = node.Next
		cnt++
	}

	removal := cnt - n

	nNode := newNode.Next

	var resNode ListNode
	resCnt := 0

	for nNode != nil {
		currVal := nNode.Val
		if resCnt != removal {
			nodeAppend(&resNode, currVal)
		}
		nNode = nNode.Next
		resCnt++
	}

	return resNode.Next
}

/* a better solution
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    start := &ListNode{ Val: 0, Next: head }
    slow, fast := start, start

    for i := 1; i <= n+1; i++ {
        fast = fast.Next
    }

    for fast != nil {
        slow = slow.Next
        fast = fast.Next
    }

    slow.Next = slow.Next.Next
    return start.Next
}
*/
