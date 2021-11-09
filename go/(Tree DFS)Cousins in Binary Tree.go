//https://leetcode.com/problems/cousins-in-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 var arr1 []int
 var arr2 []int
 func isCousins(root *TreeNode, x int, y int) bool {
	 if root.Val == x || root.Val == y { return false }
	 arr1 = make([]int, 2)
	 arr2 = make([]int, 2)
	 if root.Left != nil {
		 helper(root.Left, x, y, 1, root.Val)
	 }
	 if root.Right != nil {
		 helper(root.Right, x, y, 1, root.Val)
	 }
	 if arr1[0]!=arr2[0] && arr1[1]==arr2[1] { return true }
	 return false
 }
 
 func helper(node *TreeNode, x int, y int, depth int, parent int){
	 if node.Val == x {
		 arr1[0] = parent
		 arr1[1] = depth
	 }else if node.Val == y {
		 arr2[0] = parent
		 arr2[1] = depth
	 }
	 if arr1[0]>0 && arr2[0]>0 {
		 return
	 }else{
		 if node.Left != nil {
			 helper(node.Left, x, y, depth+1, node.Val)
		 } 
		 if node.Right != nil {
			 helper(node.Right, x, y, depth+1, node.Val)
		 }
	 }
 }