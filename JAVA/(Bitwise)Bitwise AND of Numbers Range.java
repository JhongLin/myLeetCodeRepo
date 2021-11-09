//https://leetcode.com/problems/bitwise-and-of-numbers-range/

class Solution {
    public int rangeBitwiseAnd(int left, int right) {
        if(right==left) return right;
        Stack<Integer> s1 = new Stack<>();
        int digit1 = 0;
        while(left>0){
            digit1++;
            s1.push(left%2);
            left /= 2;
        }
        Stack<Integer> s2 = new Stack<>();
        int digit2 = 0;
        while(digit2<digit1&&right>0){
            digit2++;
            s2.push(right%2);
            right /= 2;
        }
        if(right>0) return 0;
        int ans = 0;
        for(int i = digit2-1 ; i>=0 ; i--){
            Integer b1 = s1.pop();
            Integer b2 = s2.pop();
            if(b1==b2)  ans+=(b1*b2)*Math.pow(2, i);
            else    break;
        }
        return ans;
    }
}