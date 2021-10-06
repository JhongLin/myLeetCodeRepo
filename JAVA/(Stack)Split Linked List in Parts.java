//https://leetcode.com/problems/split-linked-list-in-parts/

/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode[] splitListToParts(ListNode head, int k) {
        ListNode[] ans = new ListNode[k];
        Stack<ListNode> stack = new Stack<>();
        ListNode cur = head;
        int count = 0;
        while(cur!=null){
            count++;
            stack.push(cur);
            cur=cur.next;
        }
        int[] amount = new int[k];
        if(k>=count){
            for(int i = 0 ; i<count ; i++)  amount[i]=1;
        }else{
            int quotient = count/k;
            int left = count%k;
            for(int i = 0 ; i<k ; i++)  amount[i]=quotient;
            for(int i = 0 ; i<left ; i++)   amount[i]++;
        }
        //System.out.println(Arrays.toString(amount));
        for(int i = k-1 ; i>=0 ; i--){
            if(amount[i]==0)    continue;
            ListNode one = stack.peek();
            one.next = null;
            
            for(int j = 0 ; j<amount[i] ;j++)   one = stack.pop();
            ans[i]=one;
        }
        return ans;
    }
    
}