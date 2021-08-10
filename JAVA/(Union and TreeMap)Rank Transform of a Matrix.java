class Solution {
    private int get(int x, int[] find){
        while(find[x] != x){
            find[x] = find[find[x]];
            x = find[x];
        }
        return x;
    }
    private Pair<Integer, Integer> add(int x, int y, int[] find){
        int px = get(x, find), py = get(y, find); 
        find[px] = py;
        return new Pair<>(px, py);
    }
    public int[][] matrixRankTransform(int[][] matrix) {
        int n = matrix.length, m = matrix[0].length;
        int[] rank = new int[n + m];
        
        Map<Integer, List<Pair<Integer, Integer>>> invMap = new TreeMap<>();
        for(int i = 0; i < n; i++)
            for(int j = 0; j < m; j++)
                invMap.computeIfAbsent(matrix[i][j], v -> new ArrayList<>()).add(new Pair(i, j));
        
        for(int key: invMap.keySet()){
            int[] find = new int[n + m];
            for(int i = 0; i < n+m; i++)find[i] = i;
            int[] rank2 = rank.clone();
            
            for(Pair<Integer, Integer> coor: invMap.get(key)){
                Pair<Integer, Integer> res = add(coor.getKey(), coor.getValue() + n, find);
                rank2[res.getValue()] = Math.max(rank2[res.getValue()], rank2[res.getKey()]);
            }
            
            for(Pair<Integer, Integer> coor: invMap.get(key)){
                rank[coor.getKey()] = rank[coor.getValue() + n]
                                    = matrix[coor.getKey()][coor.getValue()]
                                    = rank2[get(coor.getKey(), find)] + 1;
            }
        }
        
        return matrix;
    }
}