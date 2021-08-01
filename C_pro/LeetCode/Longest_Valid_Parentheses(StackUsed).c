//https://leetcode.com/problems/longest-valid-parentheses/

typedef struct node
{   
    int val;
    struct node *next;
}Node;

typedef struct stack
{
    Node *top;
}Stack;

void initStack(Stack *s)
{
    s->top=NULL;
}

bool isEmpty(Stack *s)
{
    return (s->top==NULL);
}

void push(Stack *s, int input)
{
    Node *new = (Node*)malloc(sizeof(Node));
    if(new==NULL)
    {
        printf("Memory Allocation Failure!\n");
        exit(1);
    }
    new->val = input;
    new->next = s->top;
    s->top = new;
}

int pop(Stack *s)
{   
    int d = s->top->val;
    Node *temp = s->top;
    s->top = s->top->next;
    free(temp);
    return d;
}

int max(int a, int b)
{
    return a>b? a : b ;
}

int longestValidParentheses(char * s){
    int sLen = strlen(s);
    if(sLen<2) return 0;
    Stack s1;
    initStack(&s1);
    int ans = 0;
    for(int i = 0 ; i<sLen ; i++ )
    {
        if(s[i]=='(')
        {
            push(&s1, i);
        }
        else
        {   
            if(isEmpty(&s1)) push(&s1, i);
            else
            {
                if(s[s1.top->val]==')') push(&s1, i);
                else
                {
                    int begin = pop(&s1);
                    ans = max(ans, isEmpty(&s1)? i+1 : i-s1.top->val);
                }
            }
        }
    }
    while(!isEmpty) pop(&s1);
    return ans;
}