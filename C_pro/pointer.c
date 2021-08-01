#include<stdio.h>
#include<stdlib.h>


int main(){
    char array[] = {5, 11, 3, 17, 2, 6};
    char *p = &array[0];

    printf("++(*p) %d \n", ++(*p));
    printf("*p++ %d \n", *p++);
    printf("*(p)++ %d \n", *(p)++);
    printf("*++p %d \n", *++p);
    printf("*(++p) %d \n", *(++p));
    printf("*(p++) %d \n", *(p++));

    return 0;
}