//https://leetcode.com/problems/perfect-squares/


class Solution {
    public int numSquares(int n) {
        int rt = (int)Math.sqrt(n)+1;
        int[] element = new int[rt];
        int[] dp = new int[n+1];
        Arrays.fill(dp, 10001);
        for(int i = 1 ; i<rt ;i++){
            element[i] = i*i;
        }
        dp[0] = 0;
        int key = 0;
        while(dp[n]==10001){
            for(int i = rt-1 ; i>0 ; i--){
                for(int j = n ; j>=element[i] ; j--){
                    //if(1+dp[j-element[i]]<dp[j])    System.out.println(Integer.toString(j)+' '+Integer.toString(1+dp[j-element[i]]));
                    if(dp[j-element[i]]==key)    dp[j] = Math.min(dp[j], 1+dp[j-element[i]]);
                }
            }
            key++;
            //System.out.println(Arrays.toString(dp));
        }
        return dp[n];
    }
}