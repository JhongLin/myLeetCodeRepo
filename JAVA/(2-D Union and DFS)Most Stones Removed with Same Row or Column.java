class Solution {
    Map<Integer, Integer> map;
    int islands;
    public int removeStones(int[][] stones) {
        map = new HashMap<>();
        islands = 0;
        /*
        if(map.putIfAbsent(2,3)==null)
            System.out.println(map.putIfAbsent(2,3));
        */
        for(int i = 0; i<stones.length; i++)    union(stones[i][0], ~stones[i][1]);
        return stones.length-islands;
    }
    private int find(int x){
        if(map.putIfAbsent(x, x)==null){
            islands++;
        }else if(map.get(x)!=x){
            map.put(x, find(map.get(x)));
        }
        return map.get(x);
    }
    private void union(int a, int b){
        a = find(a);
        b = find(b);
        if(a!=b){
            map.put(a, b);
            islands--;
        }
    }
}