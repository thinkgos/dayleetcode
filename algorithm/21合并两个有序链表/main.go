package main

type ListNode struct {
	Val  int
	next *ListNode
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
	sentry := &ListNode{}

	result := sentry

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			result.next = l1
			l1 = l1.next
		} else {
			result.next = l2
			l2 = l2.next
		}
		result = result.next
	}
	if l1 != nil {
		result.next = l1
	}
	if l2 != nil {
		result.next = l2
	}
	return sentry.next
}
