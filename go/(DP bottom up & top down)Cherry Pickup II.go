//https://leetcode.com/problems/cherry-pickup-ii/

func cherryPickup(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    dp := make([][][]int, m)
    for i := range dp {
        dp[i] = make([][]int, n)
        for j := range dp[i] {
            dp[i][j] = make([]int, n)
            for k:=range dp[i][j] {
                dp[i][j][k] = -1
            }
        }
    }
    
    //res := dfs(grid, dp, m, n, 0, 0, n-1)
    res := bottomUp(grid, dp, m, n)
    /*
    for i := range dp {
        for j := range dp[i] {
            fmt.Printf("%v\n", dp[i][j])
        }
    }*/
    return res
}
func dfs(grid [][]int, dp [][][]int, m, n, r, c1, c2 int) int {
    if m==r {return 0}
    if dp[r][c1][c2]>-1 {return dp[r][c1][c2]}
    res := 0
    for a:=-1;a<=1;a++{
        for b:=-1;b<=1;b++{
            nc1, nc2 := c1+a, c2+b
            if 0<=nc1&&nc1<n && 0<=nc2&&nc2<n {
                res = max(res, dfs(grid, dp, m, n, r+1, nc1, nc2))            
            }
        }
    }
    cherries := grid[r][c1]
    if c1!=c2 {cherries+=grid[r][c2]}
    dp[r][c1][c2] = cherries+res
    return dp[r][c1][c2]
}
func max(a, b int)int{
    if a>b{
        return a
    }else{
        return b
    }
}

func bottomUp(grid [][]int, dp [][][]int, m, n int) int {
    
    for r:=m-1 ; r>=0 ;r-- {
        for c1:=0 ; c1<=r&&c1<n ; c1++ {
            for c2:=n-1 ; c2>=n-r-1&&c2>=0 ; c2-- {
                cherries := grid[r][c1]
                if c1!=c2 {cherries+=grid[r][c2]}
                if r==m-1 {
                    dp[r][c1][c2] = cherries
                    continue
                }
                res := 0
                for a:=-1 ; a<=1 ; a++ {
                    for b:=-1 ; b<=1 ; b++ {
                        if 0<=c1+a&&c1+a<n && 0<=c2+b&&c2+b<n{
                            res = max(res, dp[r+1][c1+a][c2+b])
                        }
                    }
                }
                dp[r][c1][c2] = cherries + res
            }
        }
    }
    
    
    
    
    return dp[0][0][n-1]
}