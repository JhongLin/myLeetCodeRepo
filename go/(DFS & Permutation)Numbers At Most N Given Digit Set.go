//https://leetcode.com/problems/numbers-at-most-n-given-digit-set/

func atMostNGivenDigitSet(digits []string, n int) int {
    nc := digitCount(n) //the digit number maximum number 
    nd := len(digits) //the element amount we have
    res:= 0
    for i:=1 ; i<nc ; i++ {
        p := 1
        for j:=0 ; j<i ; j++ {
            p*=nd
        }
        res += p
    }
    res += helper(digits, n, nc, nd, nc)
    return res
}

func helper(digits []string, n , nc, nd, samePos int) int {
    if samePos-1<0 {return 1}
    res := 0
    topNumber := fetchAt(n, samePos-1)
    for _, str := range digits {
        thisNum := (int)(str[0]-'0')
        if thisNum < topNumber{
            res += pow(nd, samePos-1)
        }else if thisNum==topNumber {
            res += helper(digits, n, nc, nd, samePos-1)
            break
        }else {
            break
        }
    }
    return res
}

func digitCount(a int)int {
    res := 0
    for a>0 {
        a/=10
        res++
    }
    return res
}

func fetchAt(n, c int)int{
    n/=pow(10, c)
    n%=10
    return n
}
func pow(n, e int) int{
    if e==0 {
        return 1
    }
    res := 1
    for i:=0 ; i<e ; i++ {
        res*=n
    }
    return res
}


