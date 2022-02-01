//https://leetcode.com/problems/string-to-integer-atoi/

func myAtoi(s string) int {
    res := 0
    digit, neg := false, false
    loop: 
    for _, c := range s {
        //fmt.Printf("%c\n", c)
        if digit {
            if '0'<=c&&c<='9'{
                res = res*10 + (int)(c-'0')
            }else{
                break loop
            }
            if neg{
                if 2147483648<res {
                    return -2147483648
                }
            }else{
                if 2147483647<res {return 2147483647}
            }
            
            continue
        }
        
        if c=='+' {
            digit = true
        }else if c=='-'{
            digit = true
            neg = true
        }else if '0'<=c&&c<='9'{
            digit = true
            res += (int)(c-'0')
        }else if c==' '{
            continue
        }else{
            break loop
        }
    }
    if neg {
        return -1*res
    }
    return res
}