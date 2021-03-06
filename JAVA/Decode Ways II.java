class Solution {
    public int numDecodings(String s) {
        if(s.charAt(0)=='0')    return 0;
        
        int len = s.length();
        long M = (long)1e9+7;
        long[] dp = new long[len+1];
        Arrays.fill(dp, 0);
        
        dp[0] = 1;
        dp[1] = (s.charAt(0)=='*')? 9:1;
        for(int i = 2 ; i<=len ; i++ ){
            if(s.charAt(i-1)=='0') {
                if(s.charAt(i-2)=='1' || s.charAt(i-2)=='2'){
                    dp[i] += dp[i-2];
                } else if (s.charAt(i-2)=='*') dp[i] = dp[i-2]*2;
                else    return 0;
            } else if(s.charAt(i-1)>='1'&&s.charAt(i-1)<='9') {
                dp[i] += dp[i-1];
                if (s.charAt(i-2)=='1'||s.charAt(i-2)=='2'&&s.charAt(i-1)<='6') {
                    dp[i] += dp[i-2];
                } else if (s.charAt(i-2)=='*'){
                    dp[i] += (s.charAt(i-1)<='6')? (2*dp[i-2]) : dp[i-2];
                }
            }
            else { //s.charAt(i-1) == '*'
                    dp[i] += 9 * dp[i-1];
                    if(s.charAt(i-2)=='1') dp[i]+=9*dp[i-2];
                    else if(s.charAt(i-2)=='2') dp[i]+= 6*dp[i-2];
                    else if(s.charAt(i-2)=='*') dp[i]+= 15*dp[i-2];
            }
            dp[i]%=M;
        }
        return (int)dp[(int)len];
    }
}