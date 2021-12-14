//https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

func findDisappearedNumbers(nums []int) []int {
    n := len(nums)
    ans := make([]int, 0)
    for i:=0 ; i<n ; i++ {
        if nums[i] != i+1 {
            if nums[nums[i]-1] != nums[i] {
                swap(&nums[nums[i]-1], &nums[i])
                i--
            }
        }
    }
    for i:=0 ; i<n ; i++ {
        if nums[i] != i+1 {
            ans = append(ans, i+1)
        }
    }
    //fmt.Printf("%v\n", nums)
    
    return ans
}

func swap(a, b *int){
    *a ^= *b
    *b ^= *a
    *a ^= *b
}