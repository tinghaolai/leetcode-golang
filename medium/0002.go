package main

//Runtime: 8 ms, faster than 84.14% of Go online submissions for Add Two Numbers.
//Memory Usage: 4.4 MB, less than 97.79% of Go online submissions for Add Two Numbers.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := l1
	start := result
	carry := 0
	for l1 != nil || l2 != nil || carry == 1 {
		currentVal := 0
		if (l1 != nil) {
			currentVal += l1.Val
			l1 = l1.Next
		}

		if (l2 != nil) {
			currentVal += l2.Val
			l2 = l2.Next
		}

		currentVal += carry
		if (currentVal >= 10) {
			currentVal -= 10
			carry = 1
		} else {
			carry = 0
		}

		result.Val = currentVal
		if (result.Next == nil) {
			if (carry == 0 && l1 == nil && l2 == nil) {
				break
			}

			node := ListNode{ Val: 0, Next: nil }
			result.Next = &node
		}

		result = result.Next
	}

	return start
}

//Runtime: 4 ms, faster than 96.63% of Go online submissions for Add Two Numbers.
//Memory Usage: 4.7 MB, less than 37.04% of Go online submissions for Add Two Numbers.
func addTwoNumbers_v2(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{ Val: 0, Next: nil }
	result := head
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		if (l1 != nil) {
			carry += l1.Val
			l1 = l1.Next
		}

		if (l2 != nil) {
			carry += l2.Val
			l2 = l2.Next
		}

		result.Next = &ListNode{ Val: carry % 10, Next: nil }
		result = result.Next
		carry /= 10
	}

	return head.Next
}