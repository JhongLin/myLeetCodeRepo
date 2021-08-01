#include<stdio.h>
#include<stdlib.h>
#include<string.h>

int main(){
    int a =11;
    char i2b[100];
    itoa(a, i2b, 2);
    printf("%s\n", i2b);
    int x = 2;
    int b,c,d;
    b = a | (1 << x);
    itoa(b, i2b, 2);
    printf("set 2nd bit to 1: %s\n", i2b);
    int y = 3;
    c = a & ~(1 << y);
    itoa(c, i2b, 2);
    printf("set 3rd bit to 0: %s\n", i2b);
    int z = 2;
    d = a ^ (1 << z);
    itoa(d, i2b, 2);
    printf("Reverse 2nd bit to 0: %s\n", i2b);
    //获取的某一位的值
    //#define getbit(x,y)   ((x) >> (y)&1)
    return 0;
}