//https://leetcode.com/problems/path-sum-iii/


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func pathSum(root *TreeNode, targetSum int) int {
    var Map map[int]int
    Map = make(map[int]int)
    Map[0] = 1
    return backtrack(root, 0, targetSum, Map)
}
func backtrack(node *TreeNode, sum int, target int, Map map[int]int) int {
    if node == nil {
        return 0
    }
    sum += node.Val
    res := Map[sum-target]
    val, ok := Map[sum]
    if ok {
        Map[sum] = val+1
    }else{
        Map[sum] = 1
    }
    
    res += backtrack(node.Left, sum, target, Map) + backtrack(node.Right, sum, target, Map)
    Map[sum] = Map[sum]-1
    return res
}