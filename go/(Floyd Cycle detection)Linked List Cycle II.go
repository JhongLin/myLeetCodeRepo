//https://leetcode.com/problems/linked-list-cycle-ii/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func detectCycle(head *ListNode) *ListNode {
    var res *ListNode = nil
    if head == nil{return res}
    tor, hare := head, head
    if hare.Next!=nil && hare.Next.Next!=nil {
        tor = tor.Next
        hare = hare.Next.Next
    }else{
        return res
    }
    meet := false
    for hare.Next!=nil && hare.Next.Next!=nil { 
        //they will meet at C-r 
        //Or no cycle then meet=false
        if tor==hare {
            meet = true
            break
        }
        tor = tor.Next
        hare = hare.Next.Next
    }
    if meet {
        tor = head
        for tor!=hare{
            tor = tor.Next
            hare = hare.Next
        }
        res = tor
    }
    return res
}