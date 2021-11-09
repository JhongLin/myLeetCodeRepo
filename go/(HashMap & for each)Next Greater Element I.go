//https://leetcode.com/problems/next-greater-element-i/

func nextGreaterElement(nums1 []int, nums2 []int) []int {
    len1 := len(nums1)
    len2 := len(nums2)
    var hashMap2 map[int]int
    hashMap2 = make(map[int]int)
    for i, item := range nums2 {
        hashMap2[item] = i
    }
    ans := make([]int, len1)
    for i, _ := range nums1 {
        ans[i] = -1
    }
    for i, item := range nums1 {
        for j:=hashMap2[item]+1 ; j<len2 ; j++ {
            if nums2[j] > item {
                ans[i] = nums2[j]
                break
            }
        }
    }
    
    
    return ans
}