//https://leetcode.com/problems/minimum-absolute-difference/

func minimumAbsDifference(arr []int) [][]int {
    sort.Ints(arr)
    diff := arr[1]-arr[0]
    var ans [][]int
    ans = make([][]int, 0)
    ans = append(ans, helper(arr[0], arr[1]))
    for i:=2 ; i<len(arr) ; i++ {
        thisDiff := arr[i]-arr[i-1]
        if diff>thisDiff {
            diff = thisDiff
            ans = make([][]int, 0)
            ans = append(ans, helper(arr[i-1], arr[i]))
        }else if diff == thisDiff {
            ans = append(ans, helper(arr[i-1], arr[i]))
        }
    }
    return ans
}

func helper(a, b int) []int {
    res := make([]int, 2)
    res[0], res[1] = a, b
    return res
}