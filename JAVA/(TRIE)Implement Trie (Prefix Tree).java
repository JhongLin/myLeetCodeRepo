//https://leetcode.com/problems/implement-trie-prefix-tree/

class Trie {
    Trie[] next;
    boolean isWord;
    public Trie() {
        next = new Trie[26];
        isWord = false;
    }
    
    public void insert(String word) {
        Trie cur = this;
        for(int i = 0 ; i < word.length() ; i++){
            if(cur.next[word.charAt(i)-'a']!=null){
                cur = cur.next[word.charAt(i)-'a'];
            }else{
                cur.next[word.charAt(i)-'a'] = new Trie();
                cur = cur.next[word.charAt(i)-'a'];
            }
        }
        cur.isWord = true;
    }
    
    public boolean search(String word) {
        Trie cur = this;
        for(int i = 0 ; i < word.length() ; i++){
            if(cur.next[word.charAt(i)-'a']!=null){
                cur = cur.next[word.charAt(i)-'a'];
            }else{
                return false;
            }
        }
        if(cur.isWord)  return true;
        else return false;
    }
    
    public boolean startsWith(String prefix) {
        Trie cur = this;
        for(int i = 0 ; i < prefix.length() ; i++){
            if(cur.next[prefix.charAt(i)-'a']!=null){
                cur = cur.next[prefix.charAt(i)-'a'];
            }else{
                return false;
            }
        }
        return true;
    }
}

/**
 * Your Trie object will be instantiated and called as such:
 * Trie obj = new Trie();
 * obj.insert(word);
 * boolean param_2 = obj.search(word);
 * boolean param_3 = obj.startsWith(prefix);
 */