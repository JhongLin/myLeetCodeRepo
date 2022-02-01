//https://leetcode.com/problems/koko-eating-bananas/

// Reference from https://reurl.cc/k7EZz3
// Binary Search
func minEatingSpeed(piles []int, h int) int {
    l, r := 1, 1000000000
    for l<r {
        mid := (l+r)/2
        total := 0
        for _, val := range piles{
            total += (val+mid-1)/mid 
            // total is the sum of celi(val/mid)
        }
        if total>h{
            l=mid+1
        }else{
            r=mid
        }
    }   
    return l
}
