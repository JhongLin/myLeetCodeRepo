/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */

class Solution {
    
    
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        boolean[] find = new boolean[2];
        find[0]=false;find[1]=false;
        TreeNode target = findLR(root, p, q, find);
        //System.out.println(Arrays.toString(a));
        if(target==null){
            System.out.println(Arrays.toString(find));
            System.out.println("WH");
            return root;
        }
        return target;
    }
    private TreeNode findLR(TreeNode myself, TreeNode p, TreeNode q, boolean[] f){
        if(f[0]&&f[1])  return null;
        boolean self = false;
        TreeNode left=null, right=null;
        if(myself.val==p.val){
            //System.out.println("Find here "+p.val);
            f[0] = true;
            self = true;
        }else if(myself.val==q.val){
            //System.out.println("Find here "+q.val);
            f[1] = true;
            self = true;
        }
        
        if(!(f[0]&&f[1]))
            if(myself.left!=null)  left = findLR(myself.left, p, q, f);
        if(!(f[0]&&f[1]))
            if(myself.right!=null)  right = findLR(myself.right, p, q, f);
        
        if(!self){
            if(left==null){
                return right;
            }
            else if(right==null){
                return left;
            }
            else{
                return  myself;
            }
        } else{
            return myself;
        }
    }
}