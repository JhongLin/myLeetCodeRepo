#include<stdio.h>
#include<stdlib.h>
#include<string.h>
#include<stdbool.h>

typedef struct node
{
    int data;
    struct node *next;
} Node;

typedef struct stack{
    Node *top;
}Stack;

void stackInit(Stack* s) { 
    s->top=NULL; 
} 

bool isEmpty(Stack *s){
    return (s->top == NULL);
}

void push(Stack *s, int item) {
    Node *newNode;
    newNode = (Node *)malloc(sizeof(Node));
    if (newNode==NULL){
        printf("Memory Allocation Failure\n");
        exit(1);
    }
    newNode->data = item;
    newNode->next = s->top;
    s->top = newNode;

}

int pop(Stack* s){
    if(isEmpty(s)){
        printf("The stack is empty\n");
        return -1;
    }
    int d = s->top->data;
    Node *tmp = s->top;
    s->top = s->top->next;
    free(tmp);
    return d;
}

void list(Stack* s){
    Node* tmpNode;
    tmpNode = s->top;
    printf("Stack content: ");
    while(tmpNode!=NULL){
        printf("%d ", tmpNode->data);
        tmpNode = tmpNode->next;
    }
}

char* reverse(char* st){
    Stack s;
    stackInit(&s);
    char *scan = st;
    while(*scan!='\0'){
        push(&s, (int)*scan);
        scan++;
    }
    scan = st;
    while(*scan!='\0'){
        *scan = (char)pop(&s);
        scan++;
    }
    return st;
}

int main(int argc, char *argv[])
{
    char st[100];
    strcpy(st, "Please!");
    strcpy(st, reverse(st));

    printf("%s\n", st);
    return 0;
}