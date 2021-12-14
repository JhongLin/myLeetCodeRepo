//https://leetcode.com/problems/single-element-in-a-sorted-array/

func singleNonDuplicate(nums []int) int {
    left := 0
    n := len(nums)
    right := n-1
    for left<=right{
        mid := left + (right-left)/2
        if mid-1>=0&&nums[mid]==nums[mid-1] {
            if mid%2 == 0{
                right = mid - 1
            }else{
                left = mid + 1
            }
        }else if mid+1<n&&nums[mid]==nums[mid+1] {
            if mid%2 != 0{
                right = mid - 1
            }else{
                left = mid + 1
            }
        }else{
            return nums[mid]
        }
    }
    
    
    return nums[0]
}
