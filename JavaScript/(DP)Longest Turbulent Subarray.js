class Solution {
    public int maxTurbulenceSize(int[] arr) {
        int[] dp = new int[arr.length];
        dp[0] = 1;
        int max = 1;
        if(arr.length==1)    return dp[0];
        if(arr[0]!=arr[1]){
            dp[1] = 2;
            max = 2;
        }
        else    dp[1] = 1;
        
        for(int i = 2 ; i<arr.length ; i++){
            if(arr[i]==arr[i-1])    dp[i] = 1;
            else if(arr[i]>arr[i-1]){
                if(arr[i-1]<arr[i-2]){
                    dp[i] = dp[i-1]+1;
                } else {
                    dp[i] = 2;
                }
                
            } else { //arr[i]<arr[i-1]
                if(arr[i-1]>arr[i-2]){
                    dp[i] = dp[i-1]+1;
                } else {
                    dp[i] = 2;
                }
            }
            max = Math.max(max, dp[i]);
        }
        return max;
    }
}