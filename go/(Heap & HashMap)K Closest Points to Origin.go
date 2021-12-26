//https://leetcode.com/problems/k-closest-points-to-origin/

func kClosest(points [][]int, k int) [][]int {
    n := len(points)
    records := make(map[int][]int, n)
    pq := &myHeap{}
    heap.Init(pq)
    heapLen := 0
    for i:=0 ; i<n ; i++ {
        val := points[i][0]*points[i][0] + points[i][1]*points[i][1]
        if _, exist := records[val] ; !exist {
            //fmt.Printf("f:%d\n", i)
            records[val] = []int{i}
            heap.Push(pq, val)
            heapLen++
        }else{
            //fmt.Printf("s:%d\n", i)
            records[val] = append(records[val], i)
        }
    }
    
    ans := make([][]int, k)
    ansSerial := make([]int, k)
    cur := 0
    for k>0 {
        if heapLen==0 {break}
        val := heap.Pop(pq).(int)
        heapLen--
        index := records[val]
        //fmt.Printf("%d\n", len(index))
        for i:=0 ; i<len(index)&&k>0 ; i++ {
            ansSerial[cur] = index[i]
            cur++
            k--
        }
    }
    for i:=0 ; i<cur ; i++ {
        ans[i] = []int{points[ansSerial[i]][0], points[ansSerial[i]][1]}
    }
    /*
    heap.Push(pq, 2)
    heap.Push(pq, 1)
    fmt.Printf("%d\n", heap.Pop(pq))
    fmt.Printf("%d\n", heap.Pop(pq))
    */
    
    return ans
}


type myHeap []int
func (h myHeap) Len() int{
    return len(h)
}
func (h myHeap) Swap(i, j int){
    h[i], h[j] = h[j], h[i]
}
func (h myHeap) Less(i, j int) bool{
    return h[i] < h[j]
}
func (h *myHeap) Push(x interface{}){
    *h = append(*h, x.(int))
}
func (h *myHeap) Pop() interface{}{
    old := *h
    tmp := old[len(*h)-1]
    *h = old[0: len(*h)-1]
    return tmp
}