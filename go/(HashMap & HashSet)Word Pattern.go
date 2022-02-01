//https://leetcode.com/problems/word-pattern/

func wordPattern(pattern string, s string) bool {
    strs := strings.Split(s, " ")
    n := len(pattern)
    if len(strs)!=n {return false}
    hm := make(map[byte]string)
    hm2 := make(map[string]bool)
    for i, str := range strs {
        if content, exist := hm[pattern[i]] ; !exist {
            if _, exist := hm2[str] ; exist {return false}
            hm[pattern[i]] = str
            hm2[str] = true
        }else{
            if content!=str {return false}
        }
    }
    return true
}