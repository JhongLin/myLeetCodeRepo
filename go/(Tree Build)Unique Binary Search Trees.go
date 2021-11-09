//https://leetcode.com/problems/unique-binary-search-trees/

func numTrees(n int) int {
    get := make([]int, n+1)
    get[0] = 1
    get[1] = 1
    for i:=2 ; i<=n ; i++ {
        for j:=1 ; j<=i ; j++ {
            get[i] += get[j-1] * get[i-j]
        }
    }
    return get[n]
}
