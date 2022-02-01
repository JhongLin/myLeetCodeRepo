//https://leetcode.com/problems/maximize-distance-to-closest-person/

func maxDistToClosest(seats []int) int {
    n := len(seats)
    res := 1
    leadingZero := -1
    for i:=0 ; i<n ; i++{
        if seats[i]==1 {
            if leadingZero == -1{
                res = max(res, i)
            }else{
                var closest int = (i-leadingZero)/2
                res = max(res, closest)
            }
            leadingZero = i
        }
    }
    res = max(res, n-1-leadingZero)
    return res
}

func max(a, b int) int {
    if a>b {
        return a
    }else{
        return b
    }
}