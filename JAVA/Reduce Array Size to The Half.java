class Solution {
    public int minSetSize(int[] arr) {
        HashMap<Integer, Integer> map = new HashMap<>();
        for( int i = 0 ; i<arr.length ; i++ ) {
            Integer getn = map.get(arr[i]);
            if(getn==null){
                map.put(arr[i],1);
                //getn = map.get(3);
            } else {
                map.put(arr[i], getn+1);
            }
        }
        //Map.Entry<Integer, Integer> maxEntry = null;
        //TreeSet<Integer> ts = new TreeSet<>();
        Queue<Integer> pq = new PriorityQueue<>();
        for (Map.Entry<Integer, Integer> entry : map.entrySet())
        {
            //Integer getn = entry.getKey();
            //System.out.println(getn);
            Integer getn = entry.getValue();
            //System.out.println(getn);
            //ts.add(getn);
            pq.offer(-1*getn);
            /*
            if (maxEntry == null || entry.getValue().compareTo(maxEntry.getValue()) > 0)
            {
                maxEntry = entry;
            }
            */
        }
        int ans=0;
        int nowanumber = arr.length;
        while(nowanumber>arr.length/2){
            Integer last = -1*pq.poll();
            //System.out.println(last);
            ans++;
            nowanumber -= last;
            //last = ts.last();
            //System.out.println(last);
        }
        return ans;
    
    }
}