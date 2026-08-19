/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		// prev node is now next(reverse)
		curr.Next = prev // for first el, it will be nil
		prev = curr // update prev, in each iteration the curr of that iteration is prev for next iteration
		curr = next // incr the loop
	}
	return prev // return the reversed linked list
}
