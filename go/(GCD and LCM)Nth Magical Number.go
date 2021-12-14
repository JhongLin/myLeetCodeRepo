//https://leetcode.com/problems/nth-magical-number/

func nthMagicalNumber(n int, a int, b int) int {
    if a>b {
        temp:=a
        a = b
        b = temp
    }
    mod := 1000000007
    mul := make([]int, 0)
    L := lcm(a,b)
    
    aKey, bKey := 1, 1
    for a*aKey<L || b*bKey<L{
        if a*aKey < b*bKey  {
            mul = append(mul, a*aKey)
            aKey++
        }else{
            mul = append(mul, b*bKey)
            bKey++
        }
    }
    mul = append(mul, L)
    amount := len(mul)
    nth := (n % amount)-1
    
    n /= amount
    if nth==-1{
        nth+=amount
        n--
    }
    /*
    times := mod/L
    for n>times {
        mul[nth] = (mul[nth] + L*times)%mod
        n-=times
    }*/
    mul[nth] = (mul[nth] + L*n)%mod
    return mul[nth]
}

func gcd(a, b int) int { //a<b
    arr := make([]int, 2)
    arr[0], arr[1] = a, b
    firstKey := true
    done := false
    for !done {
        if firstKey{
            if arr[1]%arr[0] == 0{
                done = true
            }else{
                arr[1] %= arr[0]
                firstKey = false
            }
        }else{
            if arr[0]%arr[1] == 0{
                done = true
            }else{
                arr[0] %= arr[1]
                firstKey = true
            }
        }
    }
    if firstKey {return arr[0]}
    return arr[1]
}

func lcm(a, b int) int { //a<b
    fac := gcd(a,b)
    return (a*b)/fac
}