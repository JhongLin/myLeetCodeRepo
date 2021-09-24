class Solution {
    public String reverseOnlyLetters(String s) {
        char[] str = s.toCharArray();
        /*
        Stack<Character> letters = new Stack();
        Queue<Integer> indexes = new LinkedList<>();
        
        int count = 0;
        for(int i = 0 ; i<str.length ; i++){
            if((str[i]<='z'&&str[i]>='a')||(str[i]<='Z'&&str[i]>='A')){
                count++;
                letters.push(str[i]);
                indexes.offer(i);
            }
        }
        for(int i = 0 ; i<count ; i++){
            str[indexes.poll()] = letters.pop();
        }
        
        */
        int front = 0, end = str.length-1;
        while(front<end){
            while( front<end&&!((str[front]<='z'&&str[front]>='a')||(str[front]<='Z'&&str[front]>='A')) )  front++;
            while( front<end&&!((str[end]<='z'&&str[end]>='a')||(str[end]<='Z'&&str[end]>='A')) )  end--;
            if(front<end){
                char temp = str[front];
                str[front] = str[end];
                str[end] = temp;
                front++;
                end--;
            }else   break;
        }
        return new String(str);
    }
}