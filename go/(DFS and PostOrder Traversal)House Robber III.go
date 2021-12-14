//https://leetcode.com/problems/house-robber-iii/


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func rob(root *TreeNode) int {
    if root == nil {
        return 0
    }
    //take - Max amount if take the money from root
    //notTake - Max amount if NOT take monet from root
    take, notTake := postOrderTraversal(root)
    return max(take, notTake)
}

func postOrderTraversal(root *TreeNode) (int, int) {
    if root == nil {
        return 0, 0
    }
    takeL, notTakeL := postOrderTraversal(root.Left) //root.Left node's max amounts under both scenario
    takeR, notTakeR := postOrderTraversal(root.Right) //root.Right node's max amounts under both scenario
    take := root.Val + notTakeL + notTakeR //Take root node case
    notTake := max(takeL, notTakeL) + max(takeR,notTakeR) //NOT take root node case 
    return take, notTake
}


func max(a, b int) int{
    if a > b {
        return a
    } else {
        return b
    }
}