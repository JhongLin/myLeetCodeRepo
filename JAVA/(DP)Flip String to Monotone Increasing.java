class Solution {
    public int minFlipsMonoIncr(String s) {
        int len = s.length();
        char[] str = s.toCharArray();
        int[] Zero = new int[len], One = new int[len];
        
        if(str[0]=='0'){
            Zero[0] = 0;
            One[0] = 1;
        } else {
            Zero[0] = 1;
            One[0] = 0;
        }
        for(int i = 1 ; i < len ; i++ ){
            if(str[i]=='0'){
                Zero[i] = Zero[i-1];
                One[i] = Math.min(Zero[i-1], One[i-1])+1;
            } else {
                if(str[i-1]=='1')  Zero[i] = Zero[i-1]+1;
                else    Zero[i] = Zero[i-1]+1;
                One[i] = Math.min(Zero[i-1], One[i-1]);
            }
        }

        return Math.min(One[len-1], Zero[len-1]);
    }
}