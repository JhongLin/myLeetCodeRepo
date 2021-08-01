#include<stdio.h>
#include <stdlib.h>

int main(){
    /*
    char trial[100] = {"3"};
    printf("%c\n", trial[0]);
    trial[100]={};
    printf("%c\n", trial[0]);
    */
    int a[5]={1,2,3,4,5};
    printf("%d\n", *(a + 1));
    return 0;
}