//https://leetcode.com/problems/minimum-cost-to-move-chips-to-the-same-position/

func minCostToMoveChips(position []int) int {
    n := len(position)
    evenOdd := make([]int, 2)
    for i:=0 ; i<n ; i++ {
        evenOdd[position[i]%2]++
    }
    if evenOdd[0]>evenOdd[1] {
        return evenOdd[1]
    }
    return evenOdd[0]
}