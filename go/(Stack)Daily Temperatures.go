//https://leetcode.com/problems/daily-temperatures/

func dailyTemperatures(temperatures []int) []int {
    n := len(temperatures)
    stack := make([]int, 0)
    stackSize := 0
    ans := make([]int, n)
    
    //stack = stack[:len(stack)-1]
    //stack = append(stack, 7)
    for i:=0 ; i<n ; i++ {
        for stackSize>0 {
            if temperatures[stack[stackSize-1]]<temperatures[i] {
                ans[stack[stackSize-1]] = i-stack[stackSize-1]
                stack = stack[:len(stack)-1]
                stackSize--
            }else{
                break
            }
        }
        stack = append(stack, i)
        stackSize++
    }
    return ans
}