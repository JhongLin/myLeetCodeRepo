// https://leetcode.com/explore/item/3729

typedef struct heap_s {
    int capacity;
    int size;
    int* arr;
} heap_t;

heap_t* init(int);
void insert(heap_t*, int);
int top(heap_t*);
void pop(heap_t*);

// course list sorted by deadline asc
int deadlinecmp(void** a, void** b) { return (*(int**)a)[1] - (*(int**)b)[1]; }

int scheduleCourse(int** courses, int coursesRowSize, int* coursesColSizes)
{
    heap_t* heap = init(coursesRowSize);
    qsort(courses, coursesRowSize, sizeof(int*), deadlinecmp);
    int t = 0;
    for (int i = 0; i < coursesRowSize; i++) {
        // learn course i
        t += courses[i][0];
        insert(heap, courses[i][0]);
        // when time is over deadline
        if (t > courses[i][1]) {
            // remove the most long duration course
            t -= top(heap);
            pop(heap);
        }
    }

    return heap->size;
}

heap_t* init(int max)
{
    heap_t* heap = (heap_t*)malloc(sizeof(heap_t));
    heap->capacity = max;
    heap->size = 0;
    heap->arr = (int*)malloc(sizeof(int) * (max + 1));
    return heap;
}

int top(heap_t* heap) { return heap->arr[1]; }

void swap(int* a, int i, int j)
{
    a[i] = a[i] ^ a[j];
    a[j] = a[i] ^ a[j];
    a[i] = a[i] ^ a[j];
}

void insert(heap_t* heap, int val)
{
    heap->arr[++heap->size] = val;
    for (int i = heap->size; i > 1 && heap->arr[i] > heap->arr[i / 2]; i /= 2)
        swap(heap->arr, i, i / 2);
}
void pop(heap_t* heap)
{
    heap->arr[1] = heap->arr[heap->size--];
    for (int i = 1; i * 2 <= heap->size;) {
        // find smaller child to compare with the current node
        int child = heap->arr[i * 2] > heap->arr[i * 2 + 1] ? i * 2 : i * 2 + 1;
        if (heap->arr[child] > heap->arr[i]) {
            // swap if child is smaller
            swap(heap->arr, child, i);
            i = child;
        } else
            break;
    }
}