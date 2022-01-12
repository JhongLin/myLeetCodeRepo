//https://leetcode.com/problems/add-binary/

func addBinary(a string, b string) string {
    m, n := len(a), len(b)
    store := make([]byte, max(m,n)+1)
    counter := max(m,n)
    carry := false
    for i,j := m-1,n-1 ; i>=0||j>=0 ; i,j = i-1, j-1 {
        val := 0
        if carry {
            val++
            carry = false
        }
        if i>=0 && a[i]=='1' {val++}
        if j>=0 && b[j]=='1' {val++}
        if val>1 {
            carry = true
        }
        store[counter] = (byte)((val%2)+'0')
        counter--
    }
    var ans strings.Builder
    if carry {
        store[counter]='1'
        ans.Write(store[counter:])
    }else{
        ans.Write(store[counter+1:])
    }

    
    return ans.String()
}

func max(a, b int) int {
    if a > b {
        return a 
    }else{
        return b
    }
}