//https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/

func findMinArrowShots(points [][]int) int {
    sort.SliceStable(points, func(i,j int)bool{
        return points[i][1]<points[j][1]
    })
    //fmt.Printf("%v\n", points)
    res := 1
    pos := points[0][1]
    for i:=1 ; i<len(points) ; i++ {
        if points[i][0]<=pos{
            continue
        }
        res++
        pos = points[i][1]
    }
    return res
}