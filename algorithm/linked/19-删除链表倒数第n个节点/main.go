package main

func main() {

}

type ListNode struct {
	next *ListNode
}

func remove(head *ListNode, n int) *ListNode {
	sentry := &ListNode{
		next: head,
	}
	var pre *ListNode
	cur := sentry

	i := 0
	for head != nil {
		if i >= n {
			pre = cur
			cur = cur.next
		}
		head = head.next
		i++
	}
	pre.next = cur.next
	return sentry.next
}
