//https://leetcode.com/problems/kth-smallest-number-in-multiplication-table/

func min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}

func findKthNumber(n int, m int, k int) int {
    var l, r int = 1, n * m
    res := -1
    for l <= r {
        mid := l + (r - l) / 2
        cnt := 0
        for i := 1; i <= n; i++ {
            cnt += min((mid / i), m)
        } 
        if cnt >= k {
            res = mid
            r = mid - 1
        } else{
            l = mid + 1
        }
    }
    
    return res
}