class MedianFinder {
    Queue<Integer> minPQ;
    Queue<Integer> maxPQ;
    /** initialize your data structure here. */
    public MedianFinder() {
        minPQ = new PriorityQueue<>();
        maxPQ = new PriorityQueue<>((a1, a2)-> a2-a1);
    }
    
    public void addNum(int num) {
        if(minPQ.size()<=1 && maxPQ.size()==0)  minPQ.offer(num);
        else if (minPQ.peek()<num)  minPQ.offer(num);
        else  maxPQ.offer(num);
        
        if(minPQ.size()-maxPQ.size()>1)   maxPQ.offer(minPQ.remove());
        else if(maxPQ.size()-minPQ.size()>1)    minPQ.offer(maxPQ.remove());
    }
    
    public double findMedian() {
        if((minPQ.size()+maxPQ.size())%2==0)    return (minPQ.peek()+maxPQ.peek())*1.0/2;
        if(minPQ.size()>maxPQ.size())   return minPQ.peek();
        return maxPQ.peek();
    }
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * MedianFinder obj = new MedianFinder();
 * obj.addNum(num);
 * double param_2 = obj.findMedian();
 */