//https://leetcode.com/problems/linked-list-random-node/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 type Solution struct {
    size int
    arr []int
}


func Constructor(head *ListNode) Solution {
    res := Solution{size:0, arr:make([]int, 0)}
    for head!=nil{
        res.arr = append(res.arr, head.Val)
        res.size++
        head = head.Next
    }
    return res
}


func (this *Solution) GetRandom() int {
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
    return this.arr[r1.Intn(this.size)]
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */