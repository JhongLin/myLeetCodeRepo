#include<stdio.h>
#include<stdlib.h>
#include<assert.h>
typedef struct Aone{
    int i;
    union ace{
        char c;
    }ace;

}Aone;

int main(){
    Aone st;
    //st.i = 1;
    //st.ace.f = 5.0;
    //st.ace.c = 'c';
    //assert(st.ace.f == 5.0);
    int i;
    int size = sizeof(st);

    printf("%d\n", size);
    return 0;
}