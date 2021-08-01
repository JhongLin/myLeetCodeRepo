// https://leetcode.com/explore/item/3729


void swap(int **a, int**b)
{
    int *temp = *a;
    *a = *b;
    *b = temp;
}

void qs(int** courses, int front, int rear)
{
    int pivot = courses[rear][1];
    int slow = front-1;
    for(int i = front; i<rear ; i++)
    {
        if(courses[i][1]<=pivot) swap(&courses[++slow], &courses[i]);
    }
    swap(&courses[++slow], &courses[rear]);
    if(slow>front+1)    qs(courses, front, slow-1);
    if(slow<rear-1)     qs(courses, slow+1, rear);
}

int scheduleCourse(int** courses, int coursesSize, int* coursesColSize){
    qs(courses, 0, coursesSize-1);
    
    for(int i = 0 ; i<coursesSize ; i++)
        printf("%d %d\n", courses[i][0], courses[i][1]);
    int t=0, ans=0;
    for(int i = 0 ; i<coursesSize ; i++)
    {
        if(t+courses[i][0]<=courses[i][1])
        {
            t+= courses[i][0];
            ans++;
            //printf("T= %d now plus %d in constraint %d with serial %d\n", t, courses[i][0], courses[i][1], i);
        }
        else
        {
            int targetSerial=-1, maxDur=courses[i][0];
            for(int j = 0; j<i; j++)
            {
                if(courses[j][0]>maxDur)
                {
                    maxDur = courses[j][0];
                    targetSerial = j;
                }
            }
            if(maxDur>courses[i][0])
            {
                t = t-maxDur+courses[i][0];
                //printf("T= %d now minus %d with serial %d to replace %d\n", t, maxDur, targetSerial, i);
                courses[targetSerial][0] = -1;
                courses[targetSerial][1] = -1;
            }
            else
            {
                courses[i][0] = -1;
                courses[i][1] = -1;
            }
        }
    }
    return ans;
}