class Solution {
    public ListListInteger subsetsWithDup(int[] nums) {
        Arrays.sort(nums);
        
        ListListInteger append = new ArrayList();
        ListListInteger append_next;
        ListInteger temp = new ArrayList();
        append.add(temp);
        
        ListListInteger ans = new ArrayList();
        ans.addAll(append);
        
        ListInteger intList = Arrays.stream(nums).boxed().collect(Collectors.toList());
        ListInteger intList = Arrays.stream(nums).boxed().toList();
        int cur = nums[0], repeat =1;
        HashMapInteger, Integer varNum = new HashMap();
        for(int i = 1 ; i nums.length ; i++){
            if(nums[i]==nums[i-1]){
                repeat++;
            } else{
                varNum.put(nums[i-1], repeat);
                cur = nums[i];
                repeat = 1;
            }
        }
        varNum.put(cur, repeat);
        System.out.println(varNum);
        int x=0, vLen=varNum.size();
        Integer[] var = new Integer[vLen];
        Integer[] num = new Integer[vLen];

        for(Map.EntryInteger, Integer entry  varNum.entrySet()){
            var[x] = entry.getKey();
            num[x] = entry.getValue();
            if(x0) cur += num[x-1];
            else    cur = 0;
            varNum.put(entry.getKey(), x);
            x++;
        }
        System.out.println(varNum);
        System.out.println(Arrays.toString(var));
        System.out.println(Arrays.toString(num));
        =========count amount finish
        append_next = new ArrayList();
        for(int i = 0 ; ivLen ; i++){
            temp = new ArrayList();
            temp.add(var[i]);
            append_next.add(temp);
        }
        ans.addAll(append_next);
        append = append_next;
        if(nums.length==1)  return ans;
        for(int len = 2 ; lennums.length ; len++){
            append_next = new ArrayList();
            for(ListInteger base  append){
                Integer[] baseArray = base.toArray(new Integer[0]);
                Integer[] status = Arrays.copyOf(num, vLen);
                System.out.println(Arrays.toString(status));
                for(int i = 0 ; ilen-1 ; i++){
                    status[varNum.get(baseArray[i])]--;
                }
                int startPos = varNum.get(baseArray[len-2]);
                for(int i = startPos ; ivLen ; i++){
                    
                    if(status[i]0){
                        temp = (ListInteger) base.clone();
                        temp = new ArrayList(base);
                        temp.add(var[i]);
                        append_next.add(temp);
                    }
                }
                
            }
            ans.addAll(append_next);
            append = append_next;
        }
        
        
        
        ans.addAll(append);
        =====all element set
        temp = new ArrayList(nums.length); 
        for (int i  nums)  temp.add(i);
        append.add(temp);
        ans.add(temp);
        
        
        return ans;
    }
}
