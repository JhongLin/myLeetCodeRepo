//https://leetcode.com/problems/odd-even-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func oddEvenList(head *ListNode) *ListNode {
    if head==nil {return nil}
    cur := 2
    front := head
    if front.Next == nil {
        return head
    }
    rear := head.Next
    var outerHead *ListNode
    var outerTail *ListNode
    for rear!=nil {
        if cur%2==0 {
            front.Next = rear.Next
            rear.Next = nil
            if outerHead==nil {
                outerHead = rear
                outerTail = rear
            }else{
                outerTail.Next = rear
                outerTail = rear
            }
            
            rear = front.Next
        }else{
            front = rear
            rear = rear.Next
        }
        cur++
    }
    front.Next = outerHead
    
    
    return head
}