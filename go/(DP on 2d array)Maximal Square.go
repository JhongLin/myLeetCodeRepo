//https://leetcode.com/problems/maximal-square/

func maximalSquare(matrix [][]byte) int {
    m := len(matrix)
    n := len(matrix[0])

    res := 0
    for i:=0 ; i<m ; i++ {
        for j:=0 ; j<n ; j++ {
            if matrix[i][j] == '1' {
                res = 1
                break
            }
        }
        if res == 1 {
            break
        }
    }
    for i:=1 ; i<m ; i++ {
        for j:=1 ; j<n ; j++ {
            if  matrix[i][j] == '1' && matrix[i-1][j-1] > '0' {
                num := (int)(matrix[i-1][j-1]-'0')
                yCount, xCount := 0, 0
                for y:=i-1 ; y>=i-num ; y-- {
                    if matrix[y][j]=='0'{
                        break
                    }else{
                        yCount++
                    }
                }
                for x:=j-1 ; x>=j-num ; x-- {
                    if matrix[i][x]=='0'{
                        break
                    }else{
                        xCount++
                    }
                }
                extend := min(yCount, xCount)
                matrix[i][j] = (byte)(1+extend+(int)('0'))
                if res<extend+1{
                    res = extend+1
                }

                //fmt.Printf("%d\n", num)
            }
        }
    }
    /*
    for i:=0 ; i<m ; i++ {
        for j:= 0 ; j<n ; j++ {
            fmt.Printf("%c ", matrix[i][j])
        }
        fmt.Printf("\n")
    }
    */
    return res*res
}
func min(a, b int) int{
    if a<b {
        return a
    }else{
        return b
    }
}