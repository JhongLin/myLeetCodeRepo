class Solution {
    public int tribonacci(int n) {
        if(n==0)    return 0;
        else if(n<3)    return 1;
        int first = 0;
        int second = 1;
        int third = 1;
        for(int i = 3 ; i <= n ; i++ ){
            int temp = third;
            third += first + second;
            first = second;
            second = temp;
        }
        return third;
    }
}