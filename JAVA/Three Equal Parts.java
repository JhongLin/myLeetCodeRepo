class Solution {
    public int[] threeEqualParts(int[] arr) {
        int[] ans = new int[2];
        Arrays.fill(ans, -1);
        int ones=0;
        for(int i = 0;i<arr.length;i++){
            ones+=arr[i];
        }
        if(ones%3!=0)   return ans;
        else if (ones==0){
            ans[0] = 0;
            ans[1] = arr.length-1;
            return ans;
        }
        int zeroCount=0;
        int i, times=0;
        String last = new String("");
        for(i=arr.length-1 ; i>=0 ; i--){
            if(last.equals("")){
                if(arr[i]==1){
                    last += '1';
                    times++;
                }else {
                    zeroCount++;
                }
            } else{
                if(arr[i]==1){
                    last = '1'+last;
                    times++;
                }else {
                    last = '0'+last;
                }
            }
            if(times==ones/3)   break;
        }
        //System.out.println(last);
        int nz = 0;
        String str = new String("");
        times = 0;
        int first=0, second=0;
        for(i = i-1; i>=0; i--){
            if(str.equals("")){
                if(arr[i]==1){
                    str += '1';
                    times++;
                    //System.out.println(nz);
                    //System.out.println(zeroCount);
                    if(nz<zeroCount)    return ans;
                    else{
                        second = i+zeroCount+1;
                    }
                }else {
                    nz++;
                }
            } else{System.out.println("X");
                if(arr[i]==1){
                    str = '1'+str;
                    times++;
                }else {
                    str = '0'+str;
                }
            }
            if(times>=ones/3)   break;
        }
        if(!last.equals(str))    return ans;
        //System.out.println(last);
        nz=0;
        str = new String("");
        times = 0;
        
        for(i = i-1; i>=0; i--){
            if(str.equals("")){
                if(arr[i]==1){
                    str += '1';
                    times++;
                    if(nz<zeroCount)    return ans;
                    else{
                        //System.out.println(i);
                        //System.out.println(zeroCount);
                        first = i+zeroCount;
                    }
                }else {
                    nz++;
                }
            } else{
                if(arr[i]==1){
                    str = '1'+str;
                    times++;
                }else {
                    str = '0'+str;
                }
            }
            if(times>=ones/3)   break;
        }
        
        if(!last.equals(str))    return ans;
        else{
            ans[0] = first;
            ans[1] = second;
        }
        //System.out.println(last);
        //System.out.println(str);
        //System.out.println(second);
        return ans;
    }
}