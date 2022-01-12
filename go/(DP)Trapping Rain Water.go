//https://leetcode.com/problems/trapping-rain-water/

func trap(height []int) int {
    n := len(height)
    if n==1 {return 0}
    res := 0
    left, right := 0, n-1
    lMax, rMax := 0, 0
    for left<right{
        lMax = max(lMax, height[left])
        rMax = max(rMax, height[right])
        if height[left]<=height[right] {
            res += lMax - height[left]
            left++
        }else{
            res += rMax - height[right]
            right--
        }
    }
    return res
}

func max(a, b int)int{
    if a>b {
        return a
    }else{
        return b
    }
}