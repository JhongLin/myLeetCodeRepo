//https://leetcode.com/problems/sort-colors/


func sortColors(nums []int)  {
    red, white, blue := 0, 0, len(nums)-1
    for white<=blue {
        if nums[white]==0 {
            swap(&nums[red], &nums[white])
            red++
            white++
        }else if nums[white]==1 {
            white++
        }else{
            swap(&nums[white], &nums[blue])
            blue--
        }
    }
}


func swap(a, b *int){
    temp := *b
    *b = *a
    *a = temp
}