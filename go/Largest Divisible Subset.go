//https://leetcode.com/problems/largest-divisible-subset/

func largestDivisibleSubset(nums []int) []int {
    n := len(nums)
    sort.Ints(nums)
    ansGroup := make([][]int, n)
    res := 0
    max := 0
    for i:=0 ; i<n ; i++ {
        temp := make([]int, 0)
        maxLen := 0
        maxSer := -1
        for j:=i-1 ; j>=0 ; j-- {
            if nums[i]%nums[j]==0{
                if maxLen < len(ansGroup[j]){
                    maxLen =len(ansGroup[j])
                    maxSer = j
                }
            }
        }
        if maxSer != -1 {
            for j:=0 ; j<maxLen ; j++{
                temp = append(temp, ansGroup[maxSer][j])
            }
        }
        temp = append(temp, nums[i])
        if maxSer != -1 && maxLen+1 > max {
            max = maxLen + 1
            res = i
        }
        ansGroup[i] = temp
    }
    //fmt.Printf("%v\n", ansGro)
    
    return ansGroup[res]
}