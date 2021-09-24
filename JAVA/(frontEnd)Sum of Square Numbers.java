class Solution {
    public boolean judgeSquareSum(int c) {
        int r = (int)Math.sqrt(c);
        //System.out.println(r);
        //int[] repo = new int[r+1];
        /*
        for(int i = 0 ; i<=r ; i++){
            repo[i] = i*i;
        }
        */
        //System.out.println(Arrays.toString(repo));
        int front = 0, rear = r;
        while(front<=rear){
            int getSum = front*front+rear*rear;
            if(getSum>c){
                rear--;
            }else if(getSum<c){
                front++;
            }else{
                return true;
            }
        }
        return false;
    }
}