//https://leetcode.com/problems/smallest-integer-divisible-by-k/

func smallestRepunitDivByK(k int) int {
    remainder := 0
    for i:=1 ; i<=k ; i++ {
        remainder = (remainder*10 + 1)%k
        if remainder==0 {return i}
    }
    return -1
}