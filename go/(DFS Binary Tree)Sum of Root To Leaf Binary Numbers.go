//https://leetcode.com/problems/sum-of-root-to-leaf-binary-numbers/


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func sumRootToLeaf(root *TreeNode) int {
    return dfs(root, 0)
}

func dfs(node *TreeNode, preSum int) int { 
    if node.Left==nil && node.Right==nil{
        return node.Val+preSum
    }
    res := 0
    if node.Left!=nil {
        res += dfs(node.Left, (preSum+node.Val)*2)
    }
    if node.Right!=nil {
        res += dfs(node.Right, (preSum+node.Val)*2)
    }
    return res
}
