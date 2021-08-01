#include<stdio.h>
#include<stdlib.h>
#include<stdbool.h>
#include<string.h>
#include<math.h>

#define True 1
#define False 0

int brokenCalc(int X, int Y){
    if(X>Y)
        return X - Y;
}

int main(){
    int X=2, Y=3;
    printf("Output: %d\n", brokenCalc(X, Y));
    return 0;
}