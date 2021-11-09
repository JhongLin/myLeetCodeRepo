//https://leetcode.com/problems/surrounded-regions/

func solve(board [][]byte)  {
    m := len(board)
    n := len(board[0])
    
    for i:= 0 ; i<m ; i++ {
        if board[i][0]=='O'{
            board[i][0] = '0'
            helper(board, i, 0)
        }
        if board[i][n-1]=='O'{
            board[i][n-1] = '0'
            helper(board, i, n-1)
        }
    }
    for j:= 0 ; j<n ; j++ {
        if board[0][j]=='O'{
            board[0][j] = '0'
            helper(board, 0, j)
        }
        if board[m-1][j]=='O'{
            board[m-1][j] = '0'
            helper(board, m-1, j)
        }
    }
    
    for i:= 0 ; i<m ; i++{
        for j:= 0 ; j<n ; j++ {
            if board[i][j] == 'O' {
                board[i][j] = 'X'
            }else if board[i][j] == '0'{
                board[i][j] = 'O'
            }
        }
    }
}
func helper(board [][]byte, a int, b int) {
    if a-1>=0 && board[a-1][b]=='O'{
        board[a-1][b] = '0'
        helper(board, a-1, b)
    }
    if a+1<len(board) && board[a+1][b]=='O'{
        board[a+1][b] = '0'
        helper(board, a+1, b)
    }
    if b-1>=0 && board[a][b-1]=='O'{
        board[a][b-1] = '0'
        helper(board, a, b-1)
    }
    if b+1<len(board[0]) && board[a][b+1]=='O'{
        board[a][b+1] = '0'
        helper(board, a, b+1)
    }
}