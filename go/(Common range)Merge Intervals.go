//https://leetcode.com/problems/merge-intervals/

func merge(intervals [][]int) [][]int {
    n := len(intervals)
    res := make([][]int, 0)
    quickSort(intervals, 0, n-1)
    //fmt.Printf("%v\n", intervals)
    front, end := intervals[0][0], intervals[0][1]
    for i:=1 ; i<n ; i++ {
        if end>=intervals[i][0] {
            if intervals[i][1]>end {
                end = intervals[i][1]
            }
        }else{
            res = append(res, []int{front,end})
            front, end = intervals[i][0], intervals[i][1]
        }
    }
    res = append(res, []int{front,end})
    return res
}

func quickSort(arr [][]int, start, end int){
    pivot := arr[end][0]
    small := start - 1
    for i:=start ; i<end ; i++ {
        if arr[i][0]<pivot {
            small += 1
            arraySwap(arr, small, i)
        }
    }
    small += 1
    arraySwap(arr, small, end)
    if small>start+1 {quickSort(arr, start, small-1)}
    if small<end-1   {quickSort(arr, small+1, end)}
}

func arraySwap(arr [][]int, a, b int) {
    t0, t1 := arr[a][0], arr[a][1]
    arr[a][0], arr[a][1] = arr[b][0], arr[b][1]
    arr[b][0], arr[b][1] = t0, t1
}