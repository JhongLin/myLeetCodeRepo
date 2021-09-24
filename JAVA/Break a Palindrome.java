class Solution {
    public String breakPalindrome(String palindrome) {
        if(palindrome.length()==1)  return new String("");
        char[] str = palindrome.toCharArray();
        if(str.length%2==0)
        for(int i = 0 ; i<str.length ; i++){
            //System.out.println("X");
            if(str[i]!='a'){
                str[i] = 'a';
                break;
            }else if(str[i]=='a'&&i==str.length-1){
                str[i] = 'b';
            }
        }
        else
        for(int i = 0 ; i<str.length ; i++){
            //System.out.println("O");
            if(i==str.length/2) continue;
            if(str[i]!='a'){
                str[i] = 'a';
                break;
            }else if(str[i]=='a'&&i==str.length-1){
                str[i] = 'b';
            }
        }
        return new String(str);
    }
}