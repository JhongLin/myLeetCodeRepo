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
    private long maxProduct, finalSum;
    private long M;
    
    public int maxProduct(TreeNode root) {
        maxProduct = 0;
        M = (long)1e9+7;
        calSum(root);
        finalSum = root.val;
        findMax(root);
        return (int)(maxProduct%M);
    }
    private void calSum(TreeNode node){
        
        if(node.left==null&&node.right==null){
            //System.out.println(node.val);
            return;
        }
        long leftVal, rightVal ;
        
        if(node.left!=null){
            calSum(node.left);
            leftVal = (long)node.left.val;
        }else{
            leftVal = 0;
        }
        if(node.right!=null){
            calSum(node.right);
            rightVal = (long)node.right.val;
        }else{
            rightVal = 0;
        }

        node.val += leftVal + rightVal;
        //System.out.println(node.val);
    }
    private void findMax(TreeNode node){
        long leftVal, rightVal ;
        if(node.left!=null){
            leftVal = (long)node.left.val;
            maxProduct= Math.max(maxProduct, (finalSum-leftVal)*(leftVal));
            findMax(node.left);
        }
        if(node.right!=null){
            rightVal = (long)node.right.val;
            maxProduct= Math.max((finalSum-rightVal)*(rightVal), maxProduct);
            findMax(node.right);
        }
    }
}