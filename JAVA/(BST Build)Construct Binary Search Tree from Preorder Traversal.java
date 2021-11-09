//https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/

/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    private int n, serial;
    public TreeNode bstFromPreorder(int[] preorder) {
        TreeNode root = new TreeNode(preorder[0]);
        n = preorder.length;
        if(n==1)    return root;
        serial = 1;
        helper(root, preorder, root.val, root.val);
        
        return root;
    }
    private void helper(TreeNode node, int[] arr, int parentVal, int parentMax){
        while(serial<n){
            if(arr[serial]<node.val){
                TreeNode newOne = new TreeNode(arr[serial++]);
                node.left = newOne;
                helper(node.left, arr, node.val, node.val);
            }else{
                if(arr[serial]>parentMax){
                    if(node.val==parentMax){
                        TreeNode newOne = new TreeNode(arr[serial++]);
                        node.right = newOne;
                        helper(node.right, arr, node.val, node.right.val);
                    }else{
                        return;
                    }
                }else{
                    TreeNode newOne = new TreeNode(arr[serial++]);
                    node.right = newOne;
                    helper(node.right, arr, node.val, parentMax);
                }
            }
        }//end while
    }
}