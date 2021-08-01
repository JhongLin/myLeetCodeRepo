
//https://leetcode.com/problems/design-circular-queue/



typedef struct cq{
    int *values;
    int front, rear;
    int len;
} MyCircularQueue;

MyCircularQueue* myCircularQueueCreate(int k);
bool myCircularQueueEnQueue(MyCircularQueue* obj, int value);
bool myCircularQueueDeQueue(MyCircularQueue* obj);
int myCircularQueueFront(MyCircularQueue* obj);
int myCircularQueueRear(MyCircularQueue* obj);
bool myCircularQueueIsEmpty(MyCircularQueue* obj);
bool myCircularQueueIsFull(MyCircularQueue* obj);
void myCircularQueueFree(MyCircularQueue* obj);

MyCircularQueue* myCircularQueueCreate(int k) {
    MyCircularQueue *newCQ = (MyCircularQueue*)malloc(sizeof(MyCircularQueue));
    newCQ->values = (int*)calloc(k+1, sizeof(int));
    newCQ->front = 0 ; 
    newCQ->rear = 0 ;
    newCQ->len = k+1;
    return newCQ;
}

bool myCircularQueueEnQueue(MyCircularQueue* obj, int value) {
    if(myCircularQueueIsFull(obj)) return false;
    else
    {
        obj->rear = (obj->rear+1)%obj->len;
        obj->values[obj->rear] = value;
        return true;
    } 
}

bool myCircularQueueDeQueue(MyCircularQueue* obj) {
    if(myCircularQueueIsEmpty(obj)) return false;
    else
    {   
        obj->front = (obj->front+1)%obj->len;
        //int value = obj->values[front];
        return true;
    }
}

int myCircularQueueFront(MyCircularQueue* obj) {
    if(myCircularQueueIsEmpty(obj)) return -1;
    else
    {
        return obj->values[obj->front+1];
    }
}

int myCircularQueueRear(MyCircularQueue* obj) {
    if(myCircularQueueIsEmpty(obj)) return -1;
    else
    {
        return obj->values[obj->rear];
    }
}

bool myCircularQueueIsEmpty(MyCircularQueue* obj) {
    if(obj->rear==obj->front) return true;
    return false;
}

bool myCircularQueueIsFull(MyCircularQueue* obj) {
    
    if( ((obj->rear+1)%obj->len)==obj->front) return true;
    return false;
}

void myCircularQueueFree(MyCircularQueue* obj) {
    free(obj->values);
    free(obj);
}

/**
 * Your MyCircularQueue struct will be instantiated and called as such:
 * MyCircularQueue* obj = myCircularQueueCreate(k);
 * bool param_1 = myCircularQueueEnQueue(obj, value);
 
 * bool param_2 = myCircularQueueDeQueue(obj);
 
 * int param_3 = myCircularQueueFront(obj);
 
 * int param_4 = myCircularQueueRear(obj);
 
 * bool param_5 = myCircularQueueIsEmpty(obj);
 
 * bool param_6 = myCircularQueueIsFull(obj);
 
 * myCircularQueueFree(obj);
*/