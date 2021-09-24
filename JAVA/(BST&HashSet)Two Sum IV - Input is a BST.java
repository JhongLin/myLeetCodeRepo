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
    public boolean findTarget(TreeNode root, int k) {
        Queue<TreeNode> q = new LinkedList<>();
        q.offer(root);
        HashSet<Integer> hs = new HashSet<>();
        while(q.peek()!=null){
            TreeNode node = q.poll();
            if(hs.contains(k-node.val)){
                return true;
            }else{
                hs.add(node.val);
                if(node.left!=null)     q.offer(node.left);
                if(node.right!=null)    q.offer(node.right);
            }
        }
        return false;
    }
}