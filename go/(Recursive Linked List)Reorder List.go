//https://leetcode.com/problems/reorder-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reorderList(head *ListNode)  {
    helper(head, nil, head)
}
func helper(head, front, node *ListNode) *ListNode {
    if node==nil {return head}
    head = helper(head, node, node.Next)
    if head == nil {return nil}
    if head.Next!=nil && head.Next.Next!=nil {
        front.Next = node.Next
        node.Next = head.Next
        head.Next = node
        return head.Next.Next
    }else{
        return nil
    }
}