//https://leetcode.com/problems/find-the-town-judge/

func findJudge(n int, trust [][]int) int {
    trustLen := len(trust)
    if n==1&&trustLen==0 {return 1}
    trustCount := make([]int, n+1)
    trustFrom := make([]bool, n+1)
    for i:=0 ; i<trustLen ; i++ {
        trustFrom[trust[i][0]] = true
        trustCount[trust[i][1]]++
    }
    for i:=1 ; i<=n ; i++ {
        if trustFrom[i] == false && trustCount[i] == n-1 {return i}
    }
    return -1
}