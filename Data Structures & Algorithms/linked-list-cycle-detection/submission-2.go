/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    slow, fast := head, head
	for fast != nil && fast.Next != nil { // untill the faster one reaches the end(nil)
	// if nil end exists, no cycle
		slow = slow.Next // 1 step
		fast = fast.Next.Next
		if slow == fast { // if at any point fast = slwo, means fast cycled 
			return true
		}
	}
	return false
}
