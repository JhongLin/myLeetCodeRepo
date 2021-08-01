class Solution {
    public int findIntegers(int n) {
        int ans;
        /*
        List<Integer> bits = new ArrayList<>();
        for(; n>0; n>>=1){
            bits.add(n&1);
        }*/
        StringBuilder bits = new StringBuilder(Integer.toBinaryString(n)).reverse();
        //System.out.println(bits.length());
        int[] one = new int[bits.length()];
        int[] zero = new int[bits.length()];
        one[0] = zero[0] = 1;
        for(int i = 1 ; i<bits.length() ; i++){
            one[i] = zero[i-1];
            zero[i] = one[i-1] + zero[i-1];
        }
        ans = one[bits.length()-1]+zero[bits.length()-1];
        for(int i = bits.length()-2 ; i>=0 ; i--){
            if(bits.charAt(i)=='1'&&bits.charAt(i+1)=='1')  break;
            else if(bits.charAt(i)=='0'&&bits.charAt(i+1)=='0') ans-=one[i];
        }
        return ans;
    }
}