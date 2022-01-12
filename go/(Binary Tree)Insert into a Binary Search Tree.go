//https://leetcode.com/problems/insert-into-a-binary-search-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root==nil {
        res := &TreeNode{Val:val, Left:nil, Right:nil}
        return res
    }else{
        done := false
        helper(root, val, &done)
        return root
    }
}
func helper(node *TreeNode, val int, done *bool) {
    if *done {
        return 
    }
    if val<node.Val {
        if node.Left!=nil{
            helper(node.Left, val, done)
        }else{
            child := &TreeNode{Val:val, Left:nil, Right:nil}
            node.Left=child
            *done = true
        }
    }else{
        if node.Right!=nil{
            helper(node.Right, val, done)
        }else{
            child := &TreeNode{Val:val, Left:nil, Right:nil}
            node.Right=child
            *done = true
        }
    }
}