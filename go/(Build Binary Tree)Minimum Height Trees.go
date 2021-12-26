//https://leetcode.com/problems/minimum-height-trees/


func findMinHeightTrees(n int, edges [][]int) []int {
    if n == 1 {
        return []int{0}
    }
    
    m := make(map[int][]int)
    for _, v := range edges {
        m[v[0]] = append(m[v[0]], v[1])
        m[v[1]] = append(m[v[1]], v[0])
    }
    
    leaves := []int{}
    for k, v := range m {
        if len(v) == 1 {
            leaves = append(leaves, k)
        }
    }
    
    for n > 2 {
        n -= len(leaves)
        
        temp := []int{}
        for _, v := range leaves {
            childs := m[v]
            for _, c := range childs {
                m[v] = del(m[v], c)
                m[c] = del(m[c], v)
                
                if len(m[c]) == 1 {
                    temp = append(temp, c)
                }
            }
        }
        leaves = temp
    }
    return leaves
}

func del(s []int, v int) []int {
	for key, val := range s {
		if val == v {
			s = append(s[:key], s[key+1:]...)
		}
	}
	return s
}