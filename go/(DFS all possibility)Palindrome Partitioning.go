//https://leetcode.com/problems/palindrome-partitioning/

func partition(s string) [][]string {
    ans := make([][]string,0)
    mid := make([]string, 0)
    dfs(&ans, mid, s, 0)
    return ans
}

func dfs(res *[][]string, middle []string, str string, curPos int) {
    if curPos == len(str) && len(middle)>0 {
        temp := make([]string, len(middle))
        for i, content := range middle {
            temp[i] = content
        }
        *res = append(*res, temp)
    }
    for i:=curPos ; i<len(str) ; i++ {
        if isPalindrome(str, curPos, i){
            if curPos==i {
                middle = append(middle, string(str[curPos]) )
            }else{
                middle = append(middle, str[curPos:i+1])
            }
            dfs(res, middle, str, i+1)
            middle = middle[:len(middle)-1]
        }
    }
}

func isPalindrome(str string, start, end int) bool {
    for start<=end {
        if str[start]!=str[end] {return false}
        start++
        end--
    }
    return true
}