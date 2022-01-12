//https://leetcode.com/problems/populating-next-right-pointers-in-each-node/


/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

 func connect(root *Node) *Node {
    if root == nil{
        return nil
    }
    if root.Left==nil { return root }
    root.Left.Next = root.Right
    helper(root.Left, root, false)
    helper(root.Right, root, true)
	return root
}

func helper(node , parent *Node, isRight bool) {
    if isRight{
        if parent.Next != nil {
            node.Next = parent.Next.Left
        }
    }
    if node.Left!=nil { 
        node.Left.Next = node.Right
        helper(node.Left, node, false)
        helper(node.Right, node, true)
    }
}