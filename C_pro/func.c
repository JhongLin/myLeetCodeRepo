#include<stdio.h>
#include <stdlib.h>

void astrncpy(char* dst,const char* src){
    while( (*dst=*src) != '\0' ){
        dst++;
        src++;
    }
    *dst = '\0';
}

int bit_count(int num){
    int cnt = 0;
    while(num){
        num &= num - 1;
        cnt++;
    }
    return cnt;
}

int main(){
    char one[] = "Stack";
    char two[] = "Queue";
    printf("%s\n%s\n", one, two);
    astrncpy(two, one);
    printf("%s\n%s\n", one, two);
    int n = 55;
    char ns[10];
    printf("%s %d\n", itoa(n, ns, 2), bit_count(n));
    return 0;
}