//https://leetcode.com/problems/hamming-distance/


func hammingDistance(x int, y int) int {
    XOR := x ^ y
    ans := 0
    for XOR > 0 {
        ans += XOR & 1
        XOR >>= 1
    }
    return ans
}