//https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list/


/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

 func flatten(root *Node) *Node {
    if root == nil{return nil}
    helper(root)
    //check(root)
    return root
}
func helper(node *Node) *Node{
    //fmt.Printf("Now at %d \n", node.Val)
    if node.Child != nil {
        //fmt.Printf("O\n")
        if node.Next == nil {
            //fmt.Printf("OO\n")
            node.Next = node.Child
            node.Child.Prev = node
            node.Child = nil
            return helper(node.Next)
        }else{
            //fmt.Printf("X\n")
            wait := node.Next
            node.Next = node.Child
            node.Next.Prev = node
            node.Child = nil
            newNext := helper(node.Next)
            newNext.Next = wait
            wait.Prev = newNext
            return helper(wait)
        }
    }else{
        if node.Next == nil {
            return node
        } else {
            return helper(node.Next)
        }
    }
    return nil
}
func check(node *Node) {
    prev:=0
    next:=0
    if node.Prev!=nil{
        prev = node.Prev.Val
    }
    if node.Next !=nil {
        next = node.Next.Val
    }
    fmt.Printf("This Val is %d with Prev %d and Next %d\n", node.Val, prev, next)
    if node.Next!=nil{
        check(node.Next)
    }
}