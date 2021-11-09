//https://leetcode.com/problems/sort-characters-by-frequency/

func frequencySort(s string) string {
    freq := make([]int, 62)
    for _, char := range s {
        if '0'<=char && char<='9' {
            freq[char-'0']++
        } else if 'A'<=char && char<='Z' {
            freq[char-'A'+10]++
        } else { // 'a'<=char && char<='z'
            freq[char-'a'+36]++
        }
    }
    //fmt.Printf("%v", freq)
    max := -1
    var ans bytes.Buffer
    for max != 0 {
        max = 0
        serial := -1
        for i, count := range freq {
            if count>max {
                max = count
                serial = i
            }
        }
        //fmt.Printf("%d %d\n", serial, max)
        if serial==-1{
            break
        } else if 0<=serial &&serial<=9 {
            freq[serial] = 0
            for i := 0 ; i<max ; i++ {
                ans.WriteString(string('0'+serial))
            }
        }else if 10<=serial && serial<=35 {
            freq[serial] = 0
            for i := 0 ; i<max ; i++ {
                ans.WriteString(string('A'+serial-10))
            }
        } else{
            freq[serial] = 0
            for i := 0 ; i<max ; i++ {
                ans.WriteString(string('a'+serial-36))
            }
        }
    }
    return ans.String()
}