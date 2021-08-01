func suggestedProducts(products []string, searchWord string) [][]string {
    sort.Strings(products)
    //fmt.Printf("%v\n", products)
    ans := make([][]string, 0)
    //a := "tri"
    //b := "String"
    //fmt.Printf("%d\n", strings.Contains(b, a))

    for i:=0 ; i<len(searchWord) ; i++ {
        query := string(searchWord[:i+1])
        //fmt.Printf("%s\n", query)
        target := sort.Search(len(products), func(i int) bool{
            return products[i] >= query
        })
        result := make([]string, 0)
        fmt.Printf("%s %d\n", query, target)
        j, k := 0, target
        for j<3 && k<len(products) && strings.Contains(products[k][:len(query)], query){
            result = append(result, products[k])
            j++
            k++
        }
        ans = append(ans, result)
    }

    
    return ans
}