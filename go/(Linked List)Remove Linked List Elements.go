//https://leetcode.com/problems/remove-linked-list-elements/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeElements(head *ListNode, val int) *ListNode {
    if head == nil {return nil}
    var node, prev *ListNode
    prev, node = head, head.Next
    for node!=nil {
        if node.Val == val {
            node = node.Next
            prev.Next = node
        }else{
            prev = node
            node = node.Next
        }
    }
    if head.Val == val {
        head = head.Next
    }
    
    
    return head
}