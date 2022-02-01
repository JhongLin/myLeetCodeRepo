//https://leetcode.com/problems/sequential-digits/

func sequentialDigits(low int, high int) []int {
    res := make([]int, 0)
    c1, s1, h1, _ := helper(low, 0)
    c2, s2, h2, _ := helper(high, 0)
    //fmt.Println(c1, s1, h1)
    //fmt.Println(c2, s2, h2)
    init:=getSeq(0, h1, c1)
    if init<high&&s1<1&&isValid(h1, c1) {res = append(res, init)}
    for c:=c1 ; c<=c2 ; c++{
        if c==c1{
            if c==c2{
                for h:=h1+1 ; h<h2 ; h++{
                    if isValid(h, c){
                        res = append(res, getSeq(0, h, c))
                    }else{
                        break
                    }
                }
            }else{
                for h:=h1+1 ; h<10 ; h++{
                    if isValid(h, c){
                        res = append(res, getSeq(0, h, c))
                    }else{
                        break
                    }
                }
            }
        }else if c1<c&&c<c2{
            for i:=1 ; i<10 ; i++ {
                if isValid(i, c){
                    res = append(res, getSeq(0, i, c))
                }else{
                    break
                }
            }
        }else{ //c1<c && c==c2
            for i:=1 ; i<h2 ; i++ {
                if isValid(i, c){
                    res = append(res, getSeq(0, i, c))
                }else{
                    break
                }
            }
        }
    }
    if isValid(h2, c2){
        fin:=getSeq(0, h2, c2)
        if s2>-1&&fin>=low{
            if len(res)>0&&res[0]==fin{
                return res
            }else{
                res = append(res, fin)
            }
        }
    }
    return res
}
func helper(num, serial int) (int, int, int, int) {
    if num/10==0{
        return serial, 0, num, num
    }else{
        followingDigit := num%10
        digitCount, status, head, front := helper(num/10, serial+1)
        if status==0{
            if front==9||front+1>followingDigit{
                status--
            }else if front+1<followingDigit{
                status++
            }
        }
        return digitCount, status, head, followingDigit
    }
}
func isValid(h, dc int)bool { //having seq
    if h+dc<10 {
        return true
    }else{
        return false
    }
}
func getSeq(front, digit, dc int )int {
    if dc==0 {
        return front*10+digit
    }else{
        return getSeq(front*10+digit, digit+1, dc-1)
    }
}
func tenPow(n int) int {
    res:=1
    for n>0 {
        res*=10
        n--
    }
    return res
}