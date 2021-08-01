https://leetcode.com/explore/challenge/card/april-leetcoding-challenge-2021/596/week-4-april-22nd-april-28th/3718/

#define min(x,y) (x)<(y)?(x):(y);
int countBinarySubstrings(char * s){
    int len = 0;
    while(s[len]) len++;
    if(len==1) return 0;
    int curr=1, prev=0, ans = 0;
    for(int i = 1; i<len ; i++)
    {
        if(s[i]==s[i-1]) curr++;
        else
        {
            ans += min(prev, curr);
            prev = curr;
            curr = 1;
        }
    }
    return ans+min(prev, curr);
}