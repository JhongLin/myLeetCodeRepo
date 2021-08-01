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
    public TreeNode pruneTree(TreeNode root) {
        boolean zero = unZero(root);
        if(zero)    return  null;
        else    return root;
    }
    
    private boolean unZero(TreeNode node){
        if(node.val==0&&node.left==null&&node.right==null) return true;
        boolean left = false, right = false;
        if(node.left!=null) left=unZero(node.left);
        if(node.right!=null) right=unZero(node.right);
        if(left)    node.left = null;
        if(right)   node.right = null;
        if(node.val==0&&node.left==null&&node.right==null) return true;
        else    return false;
    }
}