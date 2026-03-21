package _000002_addTwoNumbers

// ListNode is structure given in problem
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 = &ListNode{Val: 0, Next: nil}
	rec(l1, l2, l3, 0)
	return l3
}

func rec(l1 *ListNode, l2 *ListNode, l3 *ListNode, carry int) {
	if l1 == nil && l2 == nil {
		if carry == 1 {
			l3.Val = carry
		}
		return
	}

	val := 0
	if l1 != nil {
		val += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		val += l2.Val
		l2 = l2.Next
	}

	l3.Val, carry = calculateValueWithCarry(val + carry)

	if l1 != nil || l2 != nil || carry == 1 {
		l3.Next = &ListNode{}
		rec(l1, l2, l3.Next, carry)
	}
}

func calculateValueWithCarry(val int) (int, int) {
	if val > 9 {
		return val - 10, 1
	}
	return val, 0
}
