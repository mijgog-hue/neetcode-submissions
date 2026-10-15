/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {
slow, fast := head, head
var prev *ListNode

for fast != nil && fast.Next != nil {
	slow = slow.Next
	fast = fast.Next.Next
}

for slow != nil {
	next := slow.Next
	slow.Next = prev
	prev = slow
	slow = next
}

for prev != nil {
	if head.Val != prev.Val {
		return false
	}
	head = head.Next
	prev = prev.Next
}
return true
}
