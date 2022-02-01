//https://leetcode.com/problems/stone-game-iv/

func winnerSquareGame(n int) bool {
    var sqrtNum int = int(math.Sqrt(float64(n)))
    sq := make([]int, sqrtNum+1)
    for i:=range sq {sq[i]=i*i}
    //fmt.Println(sq)
    dp := make([]bool, n+1)
    //fmt.Println(dp)
    for i:=1 ; i<=n ; i++{
        for j:=1;j<=sqrtNum&&sq[j]<=i;j++{
            if !dp[i-sq[j]] {
                dp[i] = true
                break
            }
        }
    }
    return dp[n]
}