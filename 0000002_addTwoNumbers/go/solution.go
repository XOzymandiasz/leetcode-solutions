package addTwoNumbers

// ListNode is structure given in problem
type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers adds two numbers represented as reversed linked lists.
// It iterates through both lists, summing corresponding digits and carry,
// building the result list in O(max(n, m)) time and O(max(n, m)) space.
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	current := &ListNode{}
	result := current
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		current.Val = sum % 10
		carry = sum / 10
		if l1 != nil || l2 != nil || carry != 0 {
			current.Next = &ListNode{}
			current = current.Next
		}
	}
	return result
}
