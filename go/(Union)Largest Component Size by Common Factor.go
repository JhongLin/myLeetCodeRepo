//https://leetcode.com/problems/largest-component-size-by-common-factor/


type DSU struct {
    p []int
}
func (d *DSU) find(x int) int {
    if d.p[x] != x {
        d.p[x] = d.find(d.p[x])
    }
    return d.p[x]
}

func (d *DSU) union(x, y int) {
    xr := d.find(x)
    yr := d.find(y)
    d.p[xr] = yr
}

func init_dsu(n int) *DSU{
    p := make([]int, n)
    for i := 0; i < n; i++ {
        p[i] = i
    }
    return &DSU{p}
}

func largestComponentSize(A []int) int {
    B := make([][]int, 0)
    for i:=0; i < len(A); i++ {
        a := A[i]
        d := 2
        facs := make([]int, 0)
        // fidn the common factor
        for d*d <= a {
            if a % d == 0 {
                for a % d == 0 {
                    a /= d
                }
                facs = append(facs, d)
            }
            d += 1
        }
        
        if a > 1 || len(facs) == 0 {
            facs = append(facs, a)
        }
        B = append(B, facs)
    }
    
    // get the unique prime factors
    hmap := make(map[int]bool)
    for _, facs := range B {
        for _, f := range facs {
            hmap[f] = true
        }
    }
    idx := 0
    primes := make([]int, 0)
    prime_to_index := make(map[int]int)
    for k, _ := range hmap {
        primes = append(primes, k)
        prime_to_index[k] = idx
        idx += 1
    }
    
    // init the DSU
    dsu := init_dsu(len(primes))
    for _, facs := range B {
        for _, f := range facs {
            dsu.union(prime_to_index[facs[0]], prime_to_index[f])
        }
    }
    
    counter := make(map[int]int)
    for _, facs := range B {
        counter[dsu.find(prime_to_index[facs[0]])] += 1
    }
    
    ans := 0
    for _, v := range counter {
        if v > ans {
            ans = v
        }
    }
    return ans
}