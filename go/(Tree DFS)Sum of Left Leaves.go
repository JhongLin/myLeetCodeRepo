//https://leetcode.com/problems/sum-of-left-leaves/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func sumOfLeftLeaves(root *TreeNode) int {
    return helper(root, false)
}
func helper(node *TreeNode, isLeft bool) int{
    sum := 0
    if node.Right==nil && node.Left==nil{
        if isLeft {
            return node.Val
        }else{
            return 0
        }
    }
    if node.Right!=nil {
        sum += helper(node.Right, false)
    }
    if node.Left!=nil {
        sum += helper(node.Left, true)
    }
    return sum
}