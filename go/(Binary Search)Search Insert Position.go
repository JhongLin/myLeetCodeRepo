//https://leetcode.com/problems/search-insert-position/


func searchInsert(nums []int, target int) int {
    n := len(nums)
    var mid, last int
    left, right := 0, n-1
    last = -1
    for mid!=last {
        last = mid
        mid = left + (right-left)/2
        if mid>=n {
            mid = n-1
            break
        }else if mid<0 {
            mid = 0
            break
        }
        if nums[mid] < target {
            left = mid + 1
        }else if nums[mid] > target {
            right = mid - 1
        }else{
            return mid
        }
        
    }
    if nums[mid]>target {
        for mid>=0 && nums[mid]>target {
            mid--
        }
        return mid+1
    }
    if nums[mid]<target {
        for mid<n && nums[mid]<target {
            mid++
        }
        return mid
    }
    return mid
}