//https://leetcode.com/problems/longest-duplicate-substring/

func longestDupSubstring(s string) string {
    ans := ""
    low := 1
    high := len(s)-1
    for low<=high {
        
        mid := low+(high-low)/2
        if len(s)==2{mid=1}
        //fmt.Printf("%d\n", mid)
        duplicate := findDup(mid, s)
        if duplicate == "" {
            high = mid -1
        }else {
            low = mid + 1
            ans = duplicate
        }

    }
    return ans
}

func findDup(strLength int, str string) string {
    hashMap := make(map[string]bool, len(str)-strLength+1)
    for i:=0 ; i<len(str)-strLength+1;i++{
        if _,exist := hashMap[str[i:i+strLength]] ; !exist {
            hashMap[str[i:i+strLength]] = true
        }else{
            return str[i:i+strLength]
        }
    }
    return ""
}