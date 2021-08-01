func isInterleave(s1 string, s2 string, s3 string) bool {
    m := len(s1)
    n := len(s2)
    if m + n != len(s3) {
        return false
    }
    
    dp := make([][]bool, 2)
    for i := 0; i < 2; i++ {
        dp[i] = make([]bool, n + 1)
    }
    
    dp[0][0] = true
    for i := 0; i <= m; i++ {
        for j := 0; j <= n; j++ {
            if i == 0 && j == 0 {
                continue
            }
            dp[i%2][j] = (i > 0 && dp[(i - 1)%2][j] && s1[i - 1] == s3[i + j - 1]) || (j > 0 && dp[i%2][j - 1] && s2[j - 1] == s3[i + j - 1])
        }
    }
    return dp[m%2][n]
}