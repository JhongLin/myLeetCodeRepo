class Solution {
    public boolean canReorderDoubled(int[] arr) {
        if(arr.length==0)   return true;
        Arrays.sort(arr);
        int nonNegSer=-1;
        for(int i = 0 ; i<arr.length ; i++){
            if(arr[i]>=0){
                nonNegSer = i;
                break;
            }
        }
        if(nonNegSer==-1)   nonNegSer = arr.length;
        Map<Integer, Integer> hash = new HashMap<>(arr.length);
        //hash.put(arr[0],1);
        for(int i = arr.length-1 ; i>=nonNegSer ; i--){
            if(hash.get(2*arr[i])!=null){
                if(hash.get(2*arr[i])==1){
                    hash.remove(2*arr[i]);
                }
                else{
                    hash.put(2*arr[i], hash.get(2*arr[i])-1);
                }
                
            } else{
                if(hash.putIfAbsent(arr[i], 1)!=null){
                    hash.put(arr[i], hash.get(arr[i])+1);
                }
            }
                                 
            //System.out.println(hash.size());
            //System.out.println(i);
        }
        
        for(int i = 0 ; i<nonNegSer ; i++){
            if(hash.get(2*arr[i])!=null){
                if(hash.get(2*arr[i])==1){
                    hash.remove(2*arr[i]);
                }
                else{
                    hash.put(2*arr[i], hash.get(2*arr[i])-1);
                }
                
            } else{
                if(hash.putIfAbsent(arr[i], 1)!=null){
                    hash.put(arr[i], hash.get(arr[i])+1);
                }
            }
        }
        if(hash.size()==0)  return true;
        else    return false;
    }
}