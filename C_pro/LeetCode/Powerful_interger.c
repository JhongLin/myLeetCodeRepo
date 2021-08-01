// https://leetcode.com/explore/item/3726



/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
void swap(int *a, int *b)
{
    int temp = *a;
    *a = *b;
    *b = temp;
}

void quickSort(int *arr, int front, int rear)
{
    int pivot = arr[rear];
    int small = front-1;
    for(int i=front; i<rear; i++)
    {
        if(arr[i]<=pivot)    swap(&arr[i], &arr[++small]);
    }
    swap(&arr[rear], &arr[++small]);
    if(small>front+1)   quickSort(arr, front, small-1);
    if(small<rear-1)    quickSort(arr, small+1, rear);
}


typedef struct node
{
    int val;
    struct node* next;
}
Node;

typedef struct stack
{
    Node *top;
}
Stack;
void push(Stack *s, int value)
{
    Node *new = (Node*)malloc(sizeof(Node));
    new->val = value;
    new->next = s->top;
    s->top = new;
}

int pop(Stack *s)
{
    if(s->top)
    {
        Node *tmp = s->top;
        int temp = tmp->val;
        s->top = s->top->next;
        free(tmp);
        return temp;
    }
    else
    {
        printf("Stack is Empty!\n");
        return -1;
    }
}

int* powerfulIntegers(int x, int y, int bound, int* returnSize){
    //printf("%d", (int)pow(2,3));
    *returnSize = 0;
    if(bound<2)
    {
        return NULL;
    }
    if(x>y)
    {
        int temp =x;
        x = y;
        y =temp;
    }
    int i, j, base=1, sum;
    Stack s1;
    s1.top = NULL;

    for(int i = 0; base<bound; )
    {
        sum = base+1;
        for(int j = 0 ; sum<=bound;)
        {
            push(&s1, sum);
            (*returnSize)++;
            j++;
            sum = base + (int)pow(y, j);
            if(y==1)    break;
        }
        i++;
        base = (int)pow(x, i);
        if(x==1)    break;
    }
    int *ans = (int*)malloc(*returnSize * sizeof(int));
    
    for(i=*returnSize-1 ; i>=0 ; i--)   ans[i] = pop(&s1);
    quickSort(ans, 0, *returnSize-1);
    int newSize = *returnSize;
    for(i=1;i<*returnSize;i++)
    {
        if(ans[i]==ans[i-1])
        {
            ans[i-1] = -1;
            newSize--;
        }
    }
    int *newAns = (int*)malloc(newSize * sizeof(int));
    for(i=0, j=0; j<*returnSize; j++)
    {
        if(ans[j]!=-1)
        {
            newAns[i] = ans[j];
            i++;
        }
    }
    *returnSize = newSize;
    return newAns;
}