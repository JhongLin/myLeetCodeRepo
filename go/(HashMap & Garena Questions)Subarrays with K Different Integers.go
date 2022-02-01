//https://leetcode.com/problems/subarrays-with-k-different-integers/

func subarraysWithKDistinct(nums []int, k int) int {
    return atMostK(nums, k)-atMostK(nums, k-1)
}

func atMostK(nums []int, k int) int {
    hm := make(map[int]int)
    left := 0
    res := 0
    //max:=0
    for right:=0 ; right<len(nums) ; right++ {
        if _, exist := hm[nums[right]] ; !exist {
            hm[nums[right]] = 1
            k--
        }else{
            if hm[nums[right]]==0 {
                k--
            }
            hm[nums[right]] += 1
        }
        for k<0 {
            hm[nums[left]]--
            if hm[nums[left]]==0 {k++}
            left++
        }
        /*
        if max<right-left+1{
            max=right-left+1
        }*/
        res += right-left+1
    }
    //fmt.Println(max)
    return res
}