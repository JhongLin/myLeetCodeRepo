func totalNQueens(n int) int {
    column := make([]int, n)
    diagonal := make([]int, n*2)
    B_diagonal := make([]int, n*2)
    var count int
    var search func(y int)
    search = func(y int){
        if y==n{
            count++
            return
        }
        for x:=0 ; x<n ; x++{
            if column[x] != 0 || diagonal[x+y]!=0 || B_diagonal[x-y+n-1]!= 0{
                continue
            }
            column[x], diagonal[x+y], B_diagonal[x-y+n-1] = 1, 1, 1
            search(y+1)
            column[x], diagonal[x+y], B_diagonal[x-y+n-1] = 0, 0, 0
        }
    }
    search(0)
    return count
}