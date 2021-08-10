class Solution {
    public int minCut(String s) {
        int len = s.length();
        int[] minCut = IntStream.range(0,len).toArray();
        for(int mid = 1; mid<len; mid++){
            for(int start=mid, end=mid; start>=0&&end<len&&s.charAt(start)==s.charAt(end); start--, end++ ){
                int newCutNum = (start==0)? 0 : minCut[start-1]+1;
                minCut[end] = Math.min(minCut[end], newCutNum);
            }
            for(int start=mid-1, end=mid; start>=0&&end<len&&s.charAt(start)==s.charAt(end); start--, end++ ){
                int newCutNum = (start==0)? 0 : minCut[start-1]+1;
                minCut[end] = Math.min(minCut[end], newCutNum);
            }
        }
        return minCut[len-1];
    }
}
//"ababdcaba"
//"addaba"