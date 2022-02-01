//https://leetcode.com/problems/detect-capital/

func detectCapitalUse(word string) bool {
    firstCap := false
    if 'A'<=word[0]&&word[0]<='Z' {
        firstCap = true
    }
    leftCap := true
    leftUnder := true
    for i:=1 ; i<len(word) ; i++ {
        if 'A'<=word[i]&&word[i]<='Z' {
            leftUnder = false
        }else{
            leftCap = false
        }
    }
    //fmt.Println(firstCap, leftCap, leftUnder)
    return (firstCap&&leftCap) || leftUnder
}