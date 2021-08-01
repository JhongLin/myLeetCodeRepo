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
    public TreeNode sortedArrayToBST(int[] nums) {
        int len = nums.length;
        int rootNum = (int)len/2;
        TreeNode root = new TreeNode(nums[rootNum]);
        if(rootNum<len-1)
            binding(root, nums, rootNum+1, len-1);
        if(rootNum>0)
            binding(root, nums, 0, rootNum-1);
        return root;
    }
    private void binding(TreeNode node, int[] nums, int begin, int end){
        int target = (begin+end)/2;
        if(node.val>nums[target]){
            node.left = new TreeNode(nums[target]);
            if(begin==end)  return;
            else{
                if(begin!=target)
                    binding(node.left, nums, begin, target-1);
                if(end!=target)
                    binding(node.left, nums, target+1, end);
            }
        }
        else{
            node.right = new TreeNode(nums[target]);
            if(begin==end)  return;
            else{
                if(begin!=target)
                    binding(node.right, nums, begin, target-1);
                if(end!=target)
                    binding(node.right, nums, target+1, end);
            }
        }
        
    }
}