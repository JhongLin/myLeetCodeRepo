//https://leetcode.com/problems/basic-calculator-ii/

func calculate(s string) int {
    operand := make([]int, 0)
    nd := 0
    operator := make([]byte, 0)
    tor := 0
    var curOperator byte
    curOperator = 'a'
    for i:=0 ; i<len(s) ; i++ {
        if s[i]==' ' {
            continue
        }else if s[i]=='+'||s[i]=='-'||s[i]=='*'||s[i]=='/'{
            if s[i]=='+'||s[i]=='-' {
                operator = append(operator, s[i])
                tor++
            }else{
                curOperator = s[i]
            }
        }else if '0'<=s[i] && s[i]<='9'{
            thisNum, newPos := getNum(s, i)
            
            i = newPos
            if curOperator!='a'{
                frontNum := operand[nd-1]
                switch curOperator {
                case '*' :
                    operand[nd-1] = frontNum*thisNum
                case '/' :
                    operand[nd-1] = frontNum/thisNum
                default:
                    fmt.Printf("Wrong\n")
                }
                curOperator = 'a'
            }else{
                operand = append(operand, thisNum)
                nd++
            }
        }
    }
    //fmt.Printf("%v\n", operand)
    //fmt.Printf("%s\n", operator)
    res := operand[0]
    
    for i:=1 ; i<nd ; i++ {
        switch operator[i-1] {
            case '+':
                res += operand[i]
            case '-':
                res -= operand[i]
            default:
                fmt.Printf("Wrong\n")
        }
    }
    return res
}

func getNum(s string, pos int) (int, int) {
    res:=0
    i:= 0
    for i=pos; i<len(s) ; i++ {
        if '0'<=s[i] && s[i]<='9' {
            res = res*10+(int)(s[i]-'0')
        }else{
            break
        }
    }
    return res, i-1
}