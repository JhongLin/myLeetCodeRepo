
class TrieNode {
    public boolean isWord;
    public int sum;
    public int origin;
    public TrieNode[] next = new TrieNode[26];
    public TrieNode() {}
}

class Trie {
    private TrieNode root;
    public Trie() {
        root = new TrieNode();
    }
    
    public void insert(String word, int val) {
        TrieNode ite = root;
        for (int i = 0; i < word.length(); ++i) {
            char c = word.charAt(i);
            if (ite.next[c - 'a'] == null) {
                ite.next[c - 'a'] = new TrieNode();
            }
            //ite.sum+=val;
            ite = ite.next[c - 'a'];
            ite.sum+=val;
        }
       // ite.sum+=val;
        ite.origin = val;
        ite.isWord = true;
    }
    
    public void update(String word, int val, int orig) {
        TrieNode ite = root;
        for (int i = 0; i < word.length(); ++i) {
            char c = word.charAt(i);
            if (ite.next[c - 'a'] == null) {
                ite.next[c - 'a'] = new TrieNode();
            }
            ite = ite.next[c - 'a'];
            ite.sum+=val;
            ite.sum-=orig;
        }
        ite.origin = val;
        ite.isWord = true;
    }
    
    public int startsSum(String prefix) {
        int ans = 0;
        TrieNode ite = root;
        for (int i = 0; i < prefix.length(); ++i) {
            char c = prefix.charAt(i);
            if (ite.next[c - 'a'] == null) {
                return 0;
            }
            ite = ite.next[c - 'a'];
            ans=ite.sum;
        }
        return ans;
    }
    
    public int search(String word) {
        TrieNode ite = root;
        int num;
        for (int i = 0; i < word.length(); ++i) {
            char c = word.charAt(i);
            if (ite.next[c - 'a'] == null) {
                return 0;
            }
            ite = ite.next[c - 'a'];
        }
        if(ite.isWord)    return ite.origin;
        else return 0;
    }
    
    public boolean startsWith(String prefix) {
        TrieNode ite = root;
        for (int i = 0; i < prefix.length(); ++i) {
            char c = prefix.charAt(i);
            if (ite.next[c - 'a'] == null) {
                return false;
            }
            ite = ite.next[c - 'a'];
        }
        return true;
    }
}



class MapSum {
    Trie trie;
    /** Initialize your data structure here. */
    public MapSum() {
        trie = new Trie();
    }
    
    public void insert(String key, int val) {
        int status = trie.search(key);
        if(status!=0){
            trie.update(key, val, status);
        }else   trie.insert(key, val);
    }
    
    public int sum(String prefix) {
        return trie.startsSum(prefix);
    }
}

/**
 * Your MapSum object will be instantiated and called as such:
 * MapSum obj = new MapSum();
 * obj.insert(key,val);
 * int param_2 = obj.sum(prefix);
 */