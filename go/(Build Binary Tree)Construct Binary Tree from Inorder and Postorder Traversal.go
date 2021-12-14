//https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func buildTree(inorder []int, postorder []int) *TreeNode {
    n := len(postorder)
    //root:= TreeNode{postorder[n-1], nil, nil}
    //root := new(TreeNode)
    //root.Val = postorder[n-1]
    //return root
    return build(inorder, postorder, 0, n-1, 0, n-1)
}
func build(inorder []int, postorder []int, i_start, i_end, p_start, p_end int) *TreeNode{
    root := new(TreeNode)
    root.Val = postorder[p_end]
    var i_pos int
    for i_pos = i_start ; i_pos<=i_end ; i_pos++ {
        if inorder[i_pos] == postorder[p_end] {
            break
        }
    }
    if i_pos>i_start{
        root.Left = build(inorder, postorder, i_start, i_pos-1, p_start, p_start+(i_pos-i_start)-1 )
    }
    if i_pos<i_end {
        root.Right = build(inorder, postorder, i_pos+1, i_end, p_start+(i_pos-i_start), p_end-1 )
    }
    return root
}