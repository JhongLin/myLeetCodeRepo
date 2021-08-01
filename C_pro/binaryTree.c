#include<stdio.h>
#include<stdlib.h>
#include<stdbool.h>
#include<assert.h>
#include<string.h>
#include<math.h>

typedef struct node{
    struct node* left;
    struct node* right;
    int value;
} Node;

typedef struct tree{
    Node *root;
} Tree;

typedef struct link{
    Node *thisNode;
    struct link *next;
} Link;

void BSTinit(Tree *t){
    t->root = NULL;
}

int Cmp(int a, int b){
    if(a<b)
        return -1;
    else if(a==b)
        return 0;
    else
        return 1;
}
void BSTinsert(Tree *t, int item){
    Node *target = t->root;
    int status;
    if(target==NULL){
        Node *newNode = (Node *)malloc(sizeof(Node));
        newNode->value = item;
        newNode->left = NULL;
        newNode->right = NULL;
        t->root = newNode;
    }
    else{
        while(1){
            status = Cmp(item, target->value);
            //printf("Insert %d with Status %d\n", item, status);
            if(status<0){
                if(target->left==NULL){
                    Node *newNode = (Node *)malloc(sizeof(Node));
                    newNode->value = item;
                    newNode->left = NULL;
                    newNode->right = NULL;
                    target->left = newNode;
                    break;
                }
                else{
                    target = target->left;
                }
            }
            else{
                if(target->right==NULL){
                    Node *newNode = (Node *)malloc(sizeof(Node));
                    newNode->value = item;
                    newNode->left = NULL;
                    newNode->right = NULL;
                    target->right = newNode;
                    break;
                }
                else{
                    target = target->right;
                }
            }
        }
    }
}

void BSTprint(Tree *t){
    Node *root = t->root;
    int n=0, LayerCount=1;
    if(root==NULL){
        puts("Tree Empty!");
        return;
    }
    Link *l = (Link *)malloc(sizeof(Link));
    Link *front, *tail;
    l->thisNode = root;
    front = l;
    tail = l;
    do{ 
        if(front->thisNode->left!=NULL){
            Link *newLink = (Link *)malloc(sizeof(Link));
            newLink->thisNode = front->thisNode->left;
            newLink->next = NULL;
            tail->next = newLink;
            tail = newLink;
            n++;
        }
        if(front->thisNode->right!=NULL){
            Link *newLink = (Link *)malloc(sizeof(Link));
            newLink->thisNode = front->thisNode->right;
            newLink->next = NULL;
            tail->next = newLink;
            tail = newLink;
            n++;
        }
        printf("%d ", front->thisNode->value);
        Link *tmp = front;
        front = front->next;
        free(tmp);
        LayerCount--;
        if(LayerCount==0){
            //printf("n: %d\n", n);
            LayerCount = n;
            puts("");
            n = 0;
        }
    }while(front!=NULL);
}

int BSTheight(Node *node){
    if(node==NULL){
        puts("Tree Empty!");
        return -1;
    }
    int left, right;
    left = node->left == NULL ? 0 : BSTheight(node->left);
    right = node->right == NULL ? 0 : BSTheight(node->right);
    return (left > right) ? (left + 1) : (right + 1);
}

int main(){
    int i;
    Tree t;
    BSTinit(&t);
    int j[] = {5, 3, 7, 2, 1, 8,9,10,11};
    /*
    for (i = 0; i < end + 1; i++){
        //printf("%d\n", i);
        BSTinsert(&t, i);
    }*/
    for (i = 0; i < (sizeof(j)/sizeof(int)); i++){
        //printf("%d\n", i);
        BSTinsert(&t, j[i]);
    }
    BSTprint(&t);
    printf("Tree Hegiht: %d\n", BSTheight(t.root)-1);
    return 0;
}