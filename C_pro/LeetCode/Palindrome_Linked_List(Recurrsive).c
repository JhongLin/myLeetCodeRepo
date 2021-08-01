//https://leetcode.com/problems/palindrome-linked-list/

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
#define True 1
#define False 0
struct ListNode* fromBack(struct ListNode* current, struct ListNode* head, bool* ans) {
    if(current==NULL)
    {
        return head;
    }
    else
    {
        head = fromBack(current->next, head, ans);
        if(*ans){
            if(head->val != current->val)
            {
                *ans = False;
                return head->next;
            }
            else
            {
                return head->next;
            }
        }
    }
    return head;
}

bool isPalindrome(struct ListNode* head){
    bool* ans = (bool*)malloc(sizeof(bool));
    *ans = True;
    fromBack(head, head, ans);
    return *ans;
}