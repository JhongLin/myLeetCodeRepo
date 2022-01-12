//https://leetcode.com/problems/number-complement/

func findComplement(num int) int {
    bits := make([]int, 32)
    cur := 0
    for num>0{
        bits[cur] = num%2
        cur++
        num/=2
    }
    res := 0
    for i:=cur-1 ; i>=0 ; i-- {
        res *= 2
        if bits[i]==0{
            res++
        }
    }
    return res
}