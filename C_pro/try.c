#include<stdio.h>
#include<stdlib.h>
#include<string.h>
#include<stdbool.h>
#include<memory.h>



int main(){
    char a[]="1133152739";
    char *b = (char *)malloc(13 * sizeof(char));
    strncpy(b, a, 6);
    printf("%d %d\n", strlen(a), strlen(b));
    puts(&a[3]);
    puts(&b[5]);

    char c[] = "abcDEF";
    char d[] = "ABCdef";
    printf("%d\n", strcmp(c, d));
    printf("%d\n", strcmp(&c[3], &d[3]));
    char e[] = "hijKLM";
    char f[] = "hijklm";
    char g[] = "gggKLM";
    printf("%d\n", strncmp(e, f, 6));
    printf("%d\n", strncmp(e, f, 3));
    printf("%d\n", strcmp(&e[3], &g[3]));
    char *new = strcat(b, e);
    printf("%s\nLen: %d\n", new, strlen(new));

    return 0;
}

/*****
 * 
 * int c;
char *dynaStringA, *dunaStringB;
char *cp;

dynaStringA = (char *)malloc(??SIZE_NUMBER??);
cp=dynaStringA;

dynaStringB = (char *)malloc(??SIZE_NUMBER??);

//取得不固定長度的使用者輸入字串
while ((c=getchar()) != '\n')
{
*cp=c;
cp++;
}
*cp='\0';

int sprintf(dynaStringB, "使用者的輸入資料是: %s", dynaStringA);
*////