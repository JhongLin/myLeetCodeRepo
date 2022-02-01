//https://leetcode.com/problems/jump-game-iv/

func minJumps(arr []int) int {
    n := len(arr)
    if n ==1 {return 0}
    
    //create new slice arr1 by removing duplicate elements from arr  
    arr1 := make([]int, 0)
    arr1 = append(arr1, arr[0])
    arr1 = append(arr1, arr[1])
    for i := 2; i < len(arr); i++ {
        if arr[i] != arr[i - 1] || arr[i - 1] != arr[i - 2] {
            arr1 = append(arr1, arr[i]) 
        }
    }
    n = len(arr1)
    
    dp := make([]int, n)
    hm := make(map[int][]int)
    for i, val := range arr1 {
        dp[i] = 50001
        if _, exist := hm[val] ; !exist{
            hm[val] = []int{i}
        }else{
            hm[val] = append(hm[val], i)
        }
    }
    dp[0] = 0
    //fmt.Println(hm)
    stack := []int{1}
    stackSize := 1
    for _,idx:=range hm[arr1[0]] {
        if idx > 1 {
            dp[idx] = 1
            stack = append(stack, idx)
            stackSize++
        }
    }
    
    for p:=2 ; p<=n ; p++ {
        if dp[n-1]<50001 {
            break
        }
        newStack := []int{}
        newStackSize := 0
        for _, idx:=range stack{
            if idx-1>=0&&dp[idx-1]>p {
                dp[idx-1] = p
                newStack = append(newStack, idx-1)
                newStackSize++
            }
            if idx+1<n&&dp[idx+1]>p {
                dp[idx+1] = p
                newStack = append(newStack, idx+1)
                newStackSize++
            }
            for _,j:=range hm[arr[idx]] {
                if j!=idx&&dp[j]>p {
                    dp[j] = p
                    newStack = append(newStack, j)
                    newStackSize++
                }
            }
            stack = newStack
            stackSize = newStackSize
        }
    }
    //fmt.Println(dp)
    return dp[n-1]
}

func min (a, b int) int {
    if a < b {
        return a
    }else{
        return b
    }
}