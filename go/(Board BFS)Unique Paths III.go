//https://leetcode.com/problems/unique-paths-iii/


func uniquePathsIII(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    var y int
    var x int
    zeros := 0
    for i:=0 ; i < m ; i++ {
        for j:=0 ; j < n ; j++ {
            if grid[i][j] == 1 {
                y = i
                x = j
            }else if grid[i][j] == 0{
                zeros++
            }
        }
    }
    return bfs(grid, y, x, m, n, 0, zeros)
}

func bfs(trails [][]int, a, b, m, n, step, zeros int) int {
    able := false
    ans := 0
    if a-1>=0 && trails[a-1][b] == 0 {
        nextTrails := copyOfSlices(trails, m, n)
        nextTrails[a-1][b] = 1
        ans += bfs(nextTrails, a-1, b, m, n, step+1, zeros)
        able = true
    }
    if a+1<m && trails[a+1][b] == 0 {
        nextTrails := copyOfSlices(trails, m, n)
        nextTrails[a+1][b] = 1
        ans += bfs(nextTrails, a+1, b, m, n, step+1, zeros)
        able = true
    }
    if b-1>=0 && trails[a][b-1] == 0 {
        nextTrails := copyOfSlices(trails, m, n)
        nextTrails[a][b-1] = 1
        ans += bfs(nextTrails, a, b-1, m, n, step+1, zeros)
        able = true
    }
    if b+1<n && trails[a][b+1] == 0 {
        nextTrails := copyOfSlices(trails, m, n)
        nextTrails[a][b+1] = 1
        ans += bfs(nextTrails, a, b+1, m, n, step+1, zeros)
        able = true
    }
    if !able{
        if (a-1>=0 && trails[a-1][b] == 2)||(a+1<m && trails[a+1][b] == 2)|| (b-1>=0 && trails[a][b-1] == 2) || (b+1<n && trails[a][b+1] == 2) {
            if step==zeros {return 1}
        }
    }
    return ans
}

func copyOfSlices(slices [][]int, m, n int) [][]int {
    newSlices := make([][]int, m)
    for i:= 0 ; i<m ; i++ {
        newSlices[i] = make([]int, n)
    }
    for i:=0 ; i < m ; i++ {
        for j:=0 ; j < n ; j++ {
            newSlices[i][j] = slices[i][j]
        }
    }
    return newSlices
}