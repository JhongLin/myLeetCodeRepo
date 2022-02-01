//https://leetcode.com/problems/gas-station/

func canCompleteCircuit(gas []int, cost []int) int {
    n := len(gas)
    start := -1
    dp, quota := make([]int, n), make([]int, n)
    totalSum := 0
    for i:=0 ; i<n ; i++{
        quota[i] = gas[i]-cost[i]
        totalSum += quota[i]
    }
    if totalSum<0 {return -1}
    dp[0] = quota[0]
    if dp[0]>=0 {start=0}

    for i:=1 ; i<n ; i++{
        if dp[i-1]<0 {
            if quota[i]>=0{
                start = i
            }
            dp[i] = quota[i]
        }else{
            dp[i] = dp[i-1]+quota[i]
            if dp[i]<0{
                start = -1
            }
        }
    }
    //fmt.Println(start, dp[n-1])
    if start==-1 {return -1}
    if start==0 {return 0}
    dp[0] = dp[n-1]+quota[0]
    if dp[0]<0 {return -1}
    for i:=1 ; i<start ; i++ {
        dp[i] = dp[i-1]+quota[i]
        if dp[i]<0 {return -1}
    }
    return start
}