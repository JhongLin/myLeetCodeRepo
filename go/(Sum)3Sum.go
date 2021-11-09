//https://leetcode.com/problems/3sum/


func threeSum(nums []int) [][]int {
    ans := make([][]int, 0)
    if len(nums)<=2 {   return ans  }
    
    //temp := []int{1,2,3}
    //ans = append(ans, temp)
    
    sort.Ints(nums)
    
    
    before_i := -100001
    for i:=0 ; i<len(nums)-2 ; i++ {
        if nums[i]==before_i {continue}
        before_j := -100001
        for j:=len(nums)-1; j>i ; j--{
            if before_j == nums[j]  {continue}
            left:=0-(nums[i]+nums[j])
            third_val := -100001
            for k:=i+1 ; k<j ; k++ {
                if nums[k]<left {
                    continue
                }else if nums[k]>left {
                    break
                }else{
                    third_val = nums[k]
                }
            }
            if third_val!=-100001 {
                temp := []int{nums[i],third_val,nums[j]}
                ans = append(ans, temp)
            }
            before_j = nums[j]
        }
        before_i = nums[i]
    }
    
    
    
    return ans
}