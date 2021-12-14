//https://leetcode.com/problems/jump-game-iii/


func canReach(arr []int, start int) bool {
    n := len(arr)
    visited := make([]bool, n)
    res := false
    helper(visited, arr, start, &res)
    //fmt.Printf("%d %v\n", n, visited)
    return res
}
func helper(visited []bool, arr []int, pos int, status *bool) {
    if *status {
        return
    }
    if arr[pos]==0{
        *status = true
        return
    }
    visited[pos] = true
    if pos-arr[pos]>=0 && !visited[pos-arr[pos]] {
        helper(visited, arr, pos-arr[pos], status)
    }
    if pos+arr[pos]<len(visited) && !visited[pos+arr[pos]] {
        helper(visited, arr, pos+arr[pos], status)
    }
}