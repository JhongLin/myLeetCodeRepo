//https://leetcode.com/problems/single-number-iii/


func singleNumber(nums []int) []int {
    ans := make([]int , 2)
    totalXOR := 0
    for _, num := range nums {
        totalXOR ^= num
    }
    rightSetBit := totalXOR & -totalXOR;
    for _, num := range nums {
        if num&rightSetBit != 0 {
            ans[0] ^= num
        }
    }
    ans[1] = ans[0]^totalXOR
    return ans
} 