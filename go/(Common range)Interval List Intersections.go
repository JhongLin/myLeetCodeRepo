//https://leetcode.com/problems/interval-list-intersections/

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
    ans := make([][]int, 0)
    i, j := 0, 0
    m, n := len(firstList), len(secondList)
    
    for i<m && j<n {
        firstList[i][0] = max(firstList[i][0], secondList[j][0])
        secondList[j][0] = firstList[i][0]
        if firstList[i][0]>firstList[i][1] {
            i++
            continue
        }else if secondList[j][0]>secondList[j][1] {
            j++
            continue
        }
        if firstList[i][1] < secondList[j][1] {
            ans = append(ans, firstList[i])
            i++
        }else if firstList[i][1] > secondList[j][1] {
            ans = append(ans, secondList[j])
            j++
        }else{
            ans = append(ans, firstList[i])
            i++
            j++
        }
        
    }
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }else {
        return b
    }
}