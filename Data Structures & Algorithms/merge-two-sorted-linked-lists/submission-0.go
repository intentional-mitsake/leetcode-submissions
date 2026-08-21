/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    d := &ListNode{} // list 
	node := d // node/ptr of the d
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			node.Next = list1
			list1 = list1.Next
		} else {
			node.Next = list2
			list2 = list2.Next
		}
		node = node.Next
	}
	// and condition so if for loop is over only one of the two lists has run out
	// the other list may have some nodes left, so this
	node.Next = list1
    if list1 == nil {
        node.Next = list2
    }


	return d.Next
}
