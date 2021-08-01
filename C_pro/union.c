#include<stdio.h>
#include<stdlib.h>
#include<assert.h>

typedef union sample_u{
    double f;
    int i[3];
    char c;
} sample_u;

typedef struct sample_s{
    double f;
    int i[3];
    char c[5];
} sample_s;

int main(){
    /*
    sample_u u1;
    u1.i = 3;
    assert(u1.i == 3);
    u1.f = 5.0;
    assert(u1.f == 5.0);
    */
    //printf("%d %.2f\n", u1.i, u1.f);
    printf("Union Size = %d\n", sizeof(sample_u));
    printf("Struct Size = %d\n", sizeof(sample_s));
    //assert(u1.i == 3);
    return 0;
}
