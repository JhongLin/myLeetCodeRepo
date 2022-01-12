//https://leetcode.com/problems/middle-of-the-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func middleNode(head *ListNode) *ListNode {
    if head.Next == nil {
        return head
    }
    front := head.Next
    rear  := head.Next.Next
    for rear!=nil && rear.Next!=nil{
        front = front.Next
        rear  = rear.Next.Next
    }
    return front
}