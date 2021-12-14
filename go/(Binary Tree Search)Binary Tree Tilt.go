//https://leetcode.com/problems/binary-tree-tilt/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func findTilt(root *TreeNode) int {
    if root==nil {return 0}
    left := 0
    right := 0
    l := 0
    r := 0
    if root.Left!=nil  {
        l = findTilt(root.Left)
        left = root.Left.Val
    }
    if root.Right!=nil {
        r = findTilt(root.Right)
        right = root.Right.Val
    }
    root.Val += left+right
    temp := 0
    if left-right<0{
        temp = right-left
    }else{
        temp = left-right
    }
    return l+r+temp
}
