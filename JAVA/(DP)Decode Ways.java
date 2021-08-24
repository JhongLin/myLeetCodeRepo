class Solution {
    public int numDecodings(String s) {
        char[] str = s.toCharArray();
        int len = str.length;
        int[] dp = new int[len+1];
        dp[0] = 1;
        if(str[0]=='0'){
            return 0;
        } else {
            dp[1] = 1;
        }
        for(int i = 2 ; i<=len ; i++){
            if(str[i-1]=='0'){
                if(str[i-2]=='1'||str[i-2]=='2'){
                    dp[i] = dp[i-2];
                } else  return 0;
            } else if(str[i-1]>'0'&&str[i-1]<='6'){
                if(str[i-2]=='1'||str[i-2]=='2')    dp[i] = dp[i-1]+dp[i-2];
                //else if (str[i-2]=='0') 
                else    dp[i] = dp[i-1];
            } else { //7. 8. 9.
                if(str[i-2]=='1')   dp[i] = dp[i-1]+dp[i-2];
                else dp[i] = dp[i-1];
            }
        }
        //System.out.println(Arrays.toString(dp));
        return dp[len];
    }
}