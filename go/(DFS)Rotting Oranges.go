//https://leetcode.com/problems/rotting-oranges/



func orangesRotting(grid [][]int) int {
    m:= len(grid)
    n:= len(grid[0])

    twos := make([][]int, 0)
    max := 0
    for i:= 0 ; i<m ; i++ {
        for j:=0 ; j<n ; j++{
            if grid[i][j]==2{
                twos = append(twos, []int{i, j})
            }
        }
    }
    
    for i:= 0 ; i<m ; i++ {
        for j:=0 ; j<n ; j++{
            if grid[i][j]>0{
                if max<grid[i][j]   {max = grid[i][j]}
                if grid[i][j] == 2 {
                    grid[i][j] = 3
                    //twos = append(twos, []int{i, j})
                    dfs(grid, i, j, -1, twos)
                }
            }
        }
    }
    if max == 0 {   
        return 0
    }else if max == 1 {
        //fmt.Printf("O\n")
        return -1
    }
    for i:= 0 ; i<m ; i++ {
        for j:=0 ; j<n ; j++{
            if grid[i][j]==1{
                return -1
            }
        }
    }
    /*
    for _, item := range grid {
            fmt.Printf("%v\n", item)
        }
    fmt.Printf("================\n")
    */
    ans := 0
    amount := len(twos)
    for amount > 0 {
        next_twos := make([][]int, 0)
        next_amount := 0
        for _, pos := range twos {
            //fmt.Printf("%v\n", twos)
            if pos[0]-1>=0 && grid[pos[0]-1][pos[1]]==-1 {
                grid[pos[0]-1][pos[1]] = 2
                next_amount++
                next_twos = append(next_twos, []int{pos[0]-1, pos[1]})
            }
            if pos[0]+1<m && grid[pos[0]+1][pos[1]]==-1 {
                grid[pos[0]+1][pos[1]] = 2
                next_amount++
                next_twos = append(next_twos, []int{pos[0]+1, pos[1]})
            }
            if pos[1]-1>=0 && grid[pos[0]][pos[1]-1]==-1 {
                grid[pos[0]][pos[1]-1] = 2
                next_amount++
                next_twos = append(next_twos, []int{pos[0], pos[1]-1})
            }
            if pos[1]+1<n && grid[pos[0]][pos[1]+1]==-1 {
                grid[pos[0]][pos[1]+1] = 2
                next_amount++
                next_twos = append(next_twos, []int{pos[0], pos[1]+1})
            }
        }
        if next_amount == 0 { break }
        /*
        for _, item := range grid {
            fmt.Printf("%v\n", item)
        }
        fmt.Printf("================\n")
        */
        ans++
        twos = next_twos
        amount = next_amount
    }
    
    
    return ans
    
}

func dfs(grid [][]int, a int, b int, serial int, repo [][]int){
    
    if a-1>=0 && grid[a-1][b] >0 {
        if grid[a-1][b] ==2 {
            //fmt.Printf("X\n")
            grid[a-1][b] = 3;
            //repo = append(repo, []int{a-1, b})
        }else if grid[a-1][b]==1&&grid[a-1][b]<=2{
            grid[a-1][b] = serial
        }
        dfs(grid, a-1, b, serial, repo)
    }
    if a+1<len(grid)&&grid[a+1][b]>0&&grid[a+1][b]<=2{
        if grid[a+1][b] ==2 {
            grid[a+1][b] = 3;
            //fmt.Printf("XX\n")
            //repo = append(repo, []int{a+1, b})
        }else if grid[a+1][b]== 1{
            grid[a+1][b] = serial
        }
        dfs(grid, a+1, b, serial, repo)
    }
    if b-1>=0 && grid[a][b-1] >0 &&grid[a][b-1]<=2{
        if grid[a][b-1] ==2 {
            //fmt.Printf("XXX\n")
            grid[a][b-1] = 3
            //repo = append(repo, []int{a, b-1})
        }else if grid[a][b-1]==1{
            grid[a][b-1] = serial
        }
        dfs(grid, a, b-1, serial, repo)
    }
    if b+1<len(grid[0])&&grid[a][b+1]>0&&grid[a][b+1]<=2{
        if grid[a][b+1] ==2 {
            //fmt.Printf("XXXX\n")
            grid[a][b+1] = 3
            //repo = append(repo, []int{a, b+1})
        }else if grid[a][b+1]==1{
            grid[a][b+1] = serial
        }
        dfs(grid, a, b+1, serial, repo)
    }
}