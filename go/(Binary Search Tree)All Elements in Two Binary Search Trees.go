//https://leetcode.com/problems/all-elements-in-two-binary-search-trees/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
    arr1 := make([]int, 0)
    if root1!=nil {helper(root1, &arr1)}
    arr2 := make([]int, 0)
    if root2!=nil {helper(root2, &arr2)}
    m, n, p1, p2 := len(arr1), len(arr2), 0, 0
    res := make([]int, m+n)
    for i:=range res {
        v1, v2 := 1000000, 1000000
        if p1<m {v1=arr1[p1]}
        if p2<n {v2=arr2[p2]}
        if v1<v2 {
            p1++
            res[i] = v1
        }else{
            p2++
            res[i] = v2
        }
    }
    return res
}
func helper(node *TreeNode, arr *[]int) {
    if node.Left!=nil {
        helper(node.Left, arr)
    }
    *arr = append(*arr, node.Val)
    if node.Right!=nil {
        helper(node.Right, arr)
    }
}