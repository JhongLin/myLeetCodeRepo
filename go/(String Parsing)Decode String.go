//https://leetcode.com/problems/decode-string/

func decodeString(s string) string {
    var res strings.Builder
    n := len(s)
    for i:=0 ; i<n ; i++ {
        if '0'<=s[i] && s[i]<='9' {
            l := i
            leftCount := 1
            var r int
            for r=i+1 ; r<n ; r++ {
                if s[r]=='['{
                    break
                }
            }
            for j:=r+1 ; j<n ; j++ {
                if s[j]=='[' {
                    leftCount++
                }else if s[j]==']' {
                    leftCount--
                }
                if leftCount == 0{
                    r = j
                    break
                }
            }
            res.Write(helper(s, l, r))
            i = r
        }else if s[i]!=']'{
            res.WriteByte(s[i])
        }
    }
    return res.String()
}

func helper(s string, left, right int) []byte {
    n := len(s)
    res := make([]byte, 0)
    repeat := 0
    for s[left]!='['{
        repeat = repeat*10 + (int)(s[left]-'0')
        left++
    }
    content := make([]byte, 0)
    for left = left+1 ; left<right ; left++ {
        if '0'<=s[left] && s[left]<='9' {
            l := left
            leftCount := 1
            var r int
            for r=left+1 ; r<n ; r++ {
                if s[r]=='['{
                    break
                }
            }
            for j:=r+1 ; j<n ; j++ {
                if s[j]=='[' {
                    leftCount++
                }else if s[j]==']' {
                    leftCount--
                }
                if leftCount == 0{
                    r = j
                    break
                }
            }
            content = append(content, helper(s, l, r)... )
            left = r
        }else if s[left]!=']' {
            content = append(content, s[left])
        }
    }
    for repeat>0 {
        res = append(res, content...)
        repeat--
    }
    return res
}