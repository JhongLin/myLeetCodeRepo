class Solution {
    public String customSortString(String order, String str) {
        if(order.equals("")||str.equals(""))  return str;
        int[] priority = new int[26];
        Arrays.fill(priority, 26);
        //char[] ord = order.toCharArray();
        char[] s   = str.toCharArray();
        //System.out.println((int)'a');
        for(int i = 0 ; i <order.length() ; i++ ) {
            priority[order.charAt(i)-'a'] = i;
        }
        /*
        for(int i = 0 ; i <26 ; i++ ) {
            System.out.println(priority[i]);
        }*/
        quickSort(s, priority, 0, s.length-1);
        //System.out.println(s);
        return new String(s);
    }
    private void quickSort(char[] str, int[] priority, int begin, int end){
        int pivot = priority[str[end]-'a'];
        int pos=begin-1;
        for(int i = begin ; i<end ; i++){
            if(priority[str[i]-'a']<pivot){
                pos++;
                charSwap(str, pos, i);
            }
        }
        pos++;
        charSwap(str, pos, end);
        if(pos-1>begin) quickSort(str, priority, begin, pos-1);
        if(pos+1<end)    quickSort(str, priority, pos+1, end);
    }
    private void charSwap(char[] str, int i, int j){
        char temp = str[j];
        str[j] = str[i];
        str[i] = temp;
    }
}