//https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/

func findMin(nums []int) int {
    if len(nums)==1 {   return nums[0]  }
    
    result := int( (^uint(0))>>1 ) //fetch int max value

    binarySearch(nums, 0, len(nums)-1, &result)
    
    return result
}

func binarySearch(arr []int, left int, right int, res *int){
    if left == right {
        if arr[left]<*res {
            *res = arr[left]
        }
        return
    }
    mid := (left+right)/2
    if arr[mid]>arr[right] {
        binarySearch(arr, mid+1, right, res)
    }else if arr[mid]<arr[right]{
        binarySearch(arr, left, mid, res)
    }else if arr[mid]==arr[right]{
        binarySearch(arr, mid+1, right, res)
        binarySearch(arr, left, mid, res)
    }
}