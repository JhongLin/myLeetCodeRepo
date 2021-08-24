class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        List<List<String>> ans = new ArrayList<>();
        HashMap<String, Integer> hash = new HashMap<>();
        int addSer = 0;
        for(int i = 0 ; i<strs.length ; i++){
            /*
            for(int j = 0 ; j<strs[i].length() ; j++){
                val+= prime[ strs[i].charAt(j)-'a' ] << strs[i].charAt(j)-'a';
            }*/
            char[] str = strs[i].toCharArray();
            Arrays.sort(str);
            String t = new String(str);
            if(hash.putIfAbsent(t, addSer)==null){
                addSer++;
                List<String> e = new ArrayList<>();
                e.add(strs[i]);
                ans.add(e);
            } else {
                ans.get(hash.get(t)).add(strs[i]);
            }
        }
        return ans;
    }
}