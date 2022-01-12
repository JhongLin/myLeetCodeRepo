//https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func maxAncestorDiff(root *TreeNode) int {
    ans := -1
    if root.Left!=nil{
        ans = helper(root.Left, root.Val, root.Val)
    }
    if root.Right!=nil{
        ans = max(ans, helper(root.Right, root.Val, root.Val))
    }
    return ans
}

func helper(node *TreeNode, fatherMax,fatherMin int) int {
    res := max(abs(node.Val-fatherMax), abs(node.Val-fatherMin))
    fatherMax = max(fatherMax, node.Val)
    fatherMin = min(fatherMin, node.Val)
    if node.Left!=nil{
        res = max(res, helper(node.Left, fatherMax, fatherMin))
    }
    if node.Right!=nil{
        res = max(res, helper(node.Right, fatherMax, fatherMin))
    }
    return res
}

func max(a, b int) int {
    if a>b{
        return a
    }else{
        return b
    }
}
func min(a,b int) int {
    if a>b {
        return b
    }else {
        return a
    }
}
func abs(a int)int{
    if a>=0{
        return a
    }else{
        return -1*a
    }
}