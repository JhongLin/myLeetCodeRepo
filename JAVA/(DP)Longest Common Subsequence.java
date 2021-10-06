//https://leetcode.com/problems/longest-common-subsequence/

class Solution {
    public int longestCommonSubsequence(String text1, String text2) {
        char[] str1 = text1.toCharArray();
        char[] str2 = text2.toCharArray();
        int ptr1, ptr2, basePtr1=0, basePtr2=0;;
        int len1=str1.length, len2=str2.length;
        int max=0;
        int[][] dp = new int[len1+1][len2+1];
        for(int i = 1; i<=len1 ;i++){
            for(int j = 1; j<=len2 ; j++){
                if(str1[i-1]==str2[j-1]){
                    dp[i][j] = dp[i-1][j-1]+1;
                }else{
                    dp[i][j] = Math.max(dp[i-1][j], dp[i][j-1]);
                }
            }
        }
        return dp[len1][len2];
    }
}