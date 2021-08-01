//https://leetcode.com/problems/longest-substring-without-repeating-characters/

#define True 1
#define False 0
#define SIZE 100


int hash(int key){
    key = key % SIZE;
    return key;
}

int lengthOfLongestSubstring(char * s){
    if(s[0]=='\0')
        return 0;
    int ascii[SIZE] = {};
    int i=0, maxLen=1, Len=0, initStart;
    while(s[i]!='\0'){
        
        int key = hash( (int)s[i] ) ;
        //printf("Now i = %d and key = %d and ascii[key] = %d\n", i, key, ascii[key]);
        if(ascii[key]!=0){
            Len = i - ascii[key]+1;
            initStart = ascii[key];
            int j;
            for (int k = 0; k < SIZE; k++)
                ascii[k] = 0;
            for (j = initStart; j <= i; j++) {
                int initKey = hash((int)s[j]);
                ascii[initKey] = j+1;
            }

        }
        else{
            ascii[key] = i + 1 ;
            Len++;
            //printf("Now i = %d and Len = %d\n", i, Len);
            if (maxLen<Len){
                maxLen = Len;
            }
        }
        i++;
        //printf("%c", s[i]);
    }


    return maxLen;
}