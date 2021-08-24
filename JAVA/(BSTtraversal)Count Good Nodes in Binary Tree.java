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
    private int count;
    public int goodNodes(TreeNode root) {
        count = 1;
        if(root.left!=null) goodCount(root.left, root.val);
        if(root.right!=null)    goodCount(root.right, root.val);
        return count;
    }
    private void goodCount(TreeNode node, int parentBiggestNum){
        if(node.val>parentBiggestNum){
            parentBiggestNum =  node.val;
            count++;
        } else if(node.val==parentBiggestNum){
            count++;
        }
        if(node.left!=null) goodCount(node.left, parentBiggestNum);
        if(node.right!=null)    goodCount(node.right, parentBiggestNum);
    }
}