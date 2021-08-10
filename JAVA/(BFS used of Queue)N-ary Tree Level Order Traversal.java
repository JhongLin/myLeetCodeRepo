/*
// Definition for a Node.
class Node {
    public int val;
    public List<Node> children;

    public Node() {}

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val, List<Node> _children) {
        val = _val;
        children = _children;
    }
};
*/

class Solution {
    public List<List<Integer>> levelOrder(Node root) {
        List<List<Integer>> ans = new ArrayList<>();
        if(root == null )   return ans;
        //System.out.println(root.children.size());
        Queue<Node> agenda = new LinkedList<>();
        agenda.offer(root);
        int layerMem = 1, nextLayerMem=0;
        List<Integer> temp = new ArrayList<>();
        while(agenda.peek()!=null){
            Node curNode = agenda.poll();
            //System.out.println(curNode.children.size());
            temp.add(curNode.val);
            nextLayerMem += curNode.children.size();
            for(int i = 0; i<curNode.children.size(); i++){
                agenda.offer(curNode.children.get(i));
            }
            layerMem--;
            if(layerMem==0){
                ans.add(temp);
                temp = new ArrayList<>();
                layerMem = nextLayerMem;
                nextLayerMem = 0;
            }
        }
        
        
        return ans;
    }
}