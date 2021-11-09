//https://leetcode.com/problems/multiply-strings/


func multiply(num1 string, num2 string) string {
    var ans strings.Builder
    len1, len2 := len(num1), len(num2)

    //fmt.Printf("%s %s\n", num1, num2)
    vals := make([]int, len1+len2-1)
    
    for i:=0 ; i<len2 ; i++ {
        base := (int)(num2[i]-'0')
        for j:=0 ; j<len1 ; j++ {
            multi := (int)(num1[j]-'0')
            vals[i+j] += base*multi
        }       
    }
    for i:=len1+len2-2 ; i>=1 ; i-- {
        //fmt.Printf("%d ", vals[i])
        vals[i-1] += vals[i]/10
        vals[i] %= 10
        //ans.WriteByte(byte('0'+vals[i]))
    }
    first := vals[0]/10
    vals[0] %= 10
    leadingZero := true
    if first>0 {
        ans.WriteByte(byte('0'+first))
        leadingZero = false
    }
    for i:=0 ; i<len1+len2-1 ; i++ {
        if leadingZero && vals[i]>0 {
            leadingZero = false
        }
        if !leadingZero {
            ans.WriteByte(byte('0'+vals[i]))
        }
    }
    if leadingZero {
        ans.WriteByte(byte('0'))
    }
    return ans.String()
}