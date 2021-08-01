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

void findMid(struct ListNode **mid, struct ListNode *head)
{
    struct ListNode  *slow = head, *fast = head;
    while(fast->next && fast->next->next)
    {
        slow = slow->next;
        fast = fast->next->next;
    }
    *mid = slow;
}

struct ListNode* reverseList(struct ListNode* head){
    struct ListNode *next=head, *prev=NULL, *cur = head;
    while(cur)
    {
        next = next->next;
        cur->next = prev;
        prev = cur;
        cur = next;
    }
    return prev;
}

void reverseLinkedList(struct ListNode *mid, struct ListNode **new)
{
    struct ListNode *next=mid, *prev=NULL, *cur=mid;
    while(cur)
    {
        next = next->next;
        cur->next = prev;
        prev = cur;
        cur = next;
    }
    *new = prev;
}

bool isPalindrome(struct ListNode *head){
    if ( head==NULL || head->next ==NULL)
    {
        return True;
    }
    struct ListNode *mid, *endHead;
    findMid(&mid, head);
    reverseLinkedList(mid, &endHead);
    
    while(head&&endHead){
        printf("%d %d\n", endHead->val, head->val);
        if( endHead->val != head->val )
        {   
            return False;
        }
        else
        {
            endHead = endHead->next;
            head = head->next;
        }
    }
    return True;
}