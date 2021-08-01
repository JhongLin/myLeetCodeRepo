class Solution {
    private int opt = Integer.MAX_VALUE;
    public List<List<String>> findLadders(String beginWord, String endWord, List<String> wordList) {
        
        List<List<String>> ans = new ArrayList<>();
        if(!wordList.contains(endWord))  return ans;
        if(wordList.contains(beginWord))    wordList.remove(beginWord);
        
        Queue<List<String>> repo = new LinkedList<>();
        List<String> temp = new ArrayList<>();
        temp.add(beginWord);
        repo.offer(temp);
        Queue<Integer> layers = new LinkedList<>();
        layers.offer(0);
        int curLayer=0;
        List<String> readyToDel = new ArrayList<>();
        while(!repo.isEmpty()){
            //System.out.println(queue.size());
            //beginWord = queue.poll();
            if(layers.peek()!=curLayer){
                
                
                for(String entry: readyToDel ){
                    wordList.remove(entry);
                }
                readyToDel = new ArrayList<>();
                
            }
            curLayer = layers.poll();
            List<String> nowFork = repo.poll();
            if(nowFork.size()>=opt){
                //System.out.println(Arrays.toString(nowFork.toArray()));
                continue;
            }
            beginWord = nowFork.get(nowFork.size()-1);
            //List<String> clone = new ArrayList<>(wordList);
            for(String entry: wordList){
                if(isTransAble(beginWord, entry)){
                    if(!entry.equals(endWord)){
                        List<String> temp2 = new ArrayList<>(nowFork);
                        temp2.add(entry);
                        repo.offer(temp2);
                        readyToDel.add(entry);
                        layers.offer(curLayer+1);
                        //clone.remove(entry);
                    }else{
                        List<String> temp2 = new ArrayList<>(nowFork);
                        temp2.add(entry);
                        if(opt>temp2.size())    opt = temp2.size();
                        ans.add(temp2);
                        //System.out.println(Arrays.toString(temp2.toArray()));
                        //System.out.println(entry);
                    }
                }
            }
            //wordList = clone;
            
        }
        for(List<String> entry : ans)
            if(entry.size()>opt)
                ans.remove(entry);
        //System.out.println(opt);
        return ans;
    }
    private boolean isTransAble(String a, String b){
        int diff=0;
        for(int i = 0 ; i<a.length() ; i++ ){
            if(a.charAt(i)!=b.charAt(i))    diff++;
        }
        if(diff==1) return true;
        return false;
    }
}