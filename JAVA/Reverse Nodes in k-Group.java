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
    public ListNode reverseKGroup(ListNode head, int k) {
        if(k==1)    return head;
        ListNode newOne = head;
        boolean firstTime=true;
        ListNode lastTail=null;
        while(newOne.next!=null){
            boolean en = true;
            ListNode temp = newOne;
            for(int i = 0 ; i < k-1 ; i++){
                if (temp.next!=null){
                    temp = temp.next;
                }else {
                    en = false;
                    break;
                }
            }
            if(en){
                ListNode prev = head, next = newOne, cur = newOne;
                for(int i = 0 ; i<k ; i++ ){
                    next = next.next;
                    cur.next = prev;
                    prev = cur;
                    cur = next;
                }
                if(firstTime){
                    firstTime = false;
                    head = prev;
                    
                }else   lastTail.next = prev;
                lastTail = newOne;
                newOne.next = cur;
                newOne = cur;
            }else{
                if(firstTime){
                    ListNode prev = null, next = newOne, cur = newOne;
                    while(cur!=null){
                        next = next.next;
                        cur.next = prev;
                        prev = cur;
                        cur = next;
                    }
                    head = prev;
                }
                break;
            }
            if(newOne==null)    break;
        }
        return head;
    }
}