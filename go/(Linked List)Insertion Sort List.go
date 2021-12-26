//https://leetcode.com/problems/insertion-sort-list/


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func insertionSortList(head *ListNode) *ListNode {
    newHead := head
    head = head.Next
    newHead.Next = nil
    for head != nil{
        node:=head
        head = head.Next
        node.Next = nil
        if node.Val<newHead.Val {
            node.Next = newHead
            newHead = node
            continue
        }
        prev := newHead
        cur  := newHead
        for cur!=nil {
            if node.Val<cur.Val {
                break
            }else{
                prev = cur
                cur  = cur.Next
            }
        }
        prev.Next = node
        node.Next = cur
    }
    return newHead
}