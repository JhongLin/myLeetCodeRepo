//https://leetcode.com/problems/can-place-flowers/

func canPlaceFlowers(flowerbed []int, n int) bool {
    if n==0 {return true}
    
    front := -1
    var count int = 0
    for i, val := range flowerbed {
        if val == 1 {
            if front==-1{
                count+=i/2
                front=i
            }else{
                count+=(i-front)/2-1
                front=i
            }
        }
        if count>=n {return true}
    }
    if front==-1{
        count += (len(flowerbed)-front)/2
    }else{
        count += (len(flowerbed)-front-1)/2
    }
    
    if count>=n {return true}
    
    return false
}