#include<stdio.h>
#include<stdlib.h>
#include<stdbool.h>
#include<string.h>
#include<math.h>

#define True 1
#define False 0
double findMedianSortedArrays(int* nums1, int nums1Size, int* nums2, int nums2Size){
    if (nums1Size == 0){
        if(nums2Size%2==0){
            int t = nums2Size / 2;
            return (double)(nums2[t - 1] + nums2[t]) / 2;
        }
        else{
            return (double)nums2[(nums2Size - 1) / 2];
        }
    }
    else if(nums2Size == 0){
        if(nums1Size%2==0){
            int t = nums1Size / 2;
            return (double)(nums1[t - 1] + nums1[t]) / 2;
        }
        else{
            return (double)nums1[(nums1Size - 1) / 2];
        }
    }
    int p1=0, p2=0;
    bool even;
    if((nums1Size+nums2Size)%2!=0){
        even = False;
    }
    else{
        even = True;
    }
    int i = 0, t = (nums1Size + nums2Size) / 2;
    double rear, front;
    if(nums1[p1]>nums2[p2]){
        front = nums2[p2];
        p2++;
    }
    else{
        front = nums1[p1];
        p1++;
    }

    
    while(i!=t){
        if(p1!=nums1Size&&p2!=nums2Size){
            if(nums1[p1]>nums2[p2]){
                rear = front;
                front = nums2[p2];
                p2++;
            }
            else{
                rear = front;
                front = nums1[p1];
                p1++;
            }
        }
        else{
            if(p1==nums1Size){
                rear = front;
                front = nums2[p2];
                p2++;
            }
            else{
                rear = front;
                front = nums1[p1];
                p1++;
            }
        }
        i++;
    }
    

    
    if(even){
        return (rear+front) / 2;
    }
    else{
        return front;
    }
}


int main(){
    int array1[] = {}, arr1_Size = 0, array2[] = {2,3}, arr2_Size = 2;
    printf("Output: %lf\n", findMedianSortedArrays(array1, arr1_Size, array2, arr2_Size));
    return 0;
}