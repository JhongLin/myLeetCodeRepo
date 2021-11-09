//https://leetcode.com/problems/insert-delete-getrandom-o1/

class RandomizedSet {
    ArrayList<Integer> repo;
    HashMap<Integer, Integer> repoLocs;
    int size;
    Random rand;
    public RandomizedSet() {
        size = 0;
        repoLocs = new HashMap<>();
        repo = new ArrayList<>();
        rand = new Random();
    }
    
    public boolean insert(int val) {
        if(repo.contains(val)){
            return false;
        }
        repoLocs.put(val, size);
        size++;
        repo.add(val);
        return true;
    }
    
    public boolean remove(int val) {
        if(!repo.contains(val)) return false;
        if(repo.get(size-1)!=val){
            int lastElement = repo.get(size-1);
            int pos = repoLocs.get(val);
            repo.set(pos, lastElement);
            repoLocs.put(lastElement, pos);
        }
        repo.remove(size-1);
        size--;
        repoLocs.remove(val);
        return true;
    }
    
    public int getRandom() {
        return repo.get( rand.nextInt(size) );
    }
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * RandomizedSet obj = new RandomizedSet();
 * boolean param_1 = obj.insert(val);
 * boolean param_2 = obj.remove(val);
 * int param_3 = obj.getRandom();
 */