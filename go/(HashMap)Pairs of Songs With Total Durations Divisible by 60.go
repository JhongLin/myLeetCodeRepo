//https://leetcode.com/problems/pairs-of-songs-with-total-durations-divisible-by-60/

func numPairsDivisibleBy60(time []int) int {
    res := 0
    remainder := make(map[int]int)
    for _, val := range time {
        tmp := val % 60
        if count, exist := remainder[(60-tmp)%60] ; exist {
            res += count
        }
        if count, exist := remainder[tmp] ; !exist {
            remainder[tmp] = 1
        }else{
            remainder[tmp] = count + 1
        }
    }
    
    
    
    
    
    return res
}