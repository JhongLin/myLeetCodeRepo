//https://leetcode.com/problems/complement-of-base-10-integer/

func bitwiseComplement(n int) int {
    if n==0 {return 1}
    res := 0
    mul := 1
    for n>0 {
        tmp := n%2
        if tmp==0 {res += mul}
        mul *= 2
        n /= 2
    }
    return res
}