//https://leetcode.com/problems/word-search-ii/

class Trie {
    Trie[] next;
    boolean isWord;
    String thisWord;
    public Trie(){
        next = new Trie[26];
        isWord = false;
        thisWord = null;
    }
    public void insert(String word){
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
        cur.thisWord = word;
    }
    public void delete(String word){
        Trie cur = this;
        Stack<Trie> stack = new Stack<>();
        for(int i = 0 ; i < word.length() ; i++){
            if(cur.next[word.charAt(i)-'a']!=null){
                cur = cur.next[word.charAt(i)-'a'];
                stack.push(cur);
            }else{
                System.out.println("no this word");
                return;
            }
        }
        for(int i = 0 ; i < word.length() ; i++){
            Trie mine = stack.pop();
            
        }
    }
}

class Solution {
    int m ;
    int n ;
    List<String> ans ;
    public List<String> findWords(char[][] board, String[] words) {
        Trie t = new Trie();
        for(int i = 0; i<words.length ; i++){
            t.insert(words[i]);
        }
        ans = new ArrayList<>();
        m = board.length;
        n = board[0].length;
        for(int i = 0 ; i<m ; i++){
            for(int j = 0 ; j<n ; j++){
                if(t.next[board[i][j]-'a']!=null){
                    //System.out.println(Integer.toString(i)+' '+Integer.toString(j));
                    char temp = board[i][j];
                    Trie why = t.next[board[i][j]-'a'];
                    board[i][j] = '0';
                    dfs(board, i, j, why);
                    board[i][j] = temp;
                    
                }
            }
        }
        
        return ans;
    }
    private void dfs(char[][] board, int row, int col, Trie t){
        if(t.isWord){
            //System.out.println("O");
            ans.add(t.thisWord);
            t.isWord = false;
            //return ;
        }
        if(row-1>=0&&board[row-1][col]!='0'&&t.next[board[row-1][col]-'a']!=null){
            char temp = board[row-1][col];
            Trie wh = t.next[board[row-1][col]-'a'];
            board[row-1][col] = '0';
            
            dfs(board, row-1, col, wh);
            //if(dfs(board, row-1, col, t))   return true;
            board[row-1][col] = temp;
            //System.out.println("X");
        }
        if(row+1<m&&board[row+1][col]!='0'&&t.next[board[row+1][col]-'a']!=null){
            char temp = board[row+1][col];
            Trie wh = t.next[board[row+1][col]-'a'];
            board[row+1][col] = '0';
            dfs(board, row+1, col, wh);
            //if(dfs(board, row-1, col, t))   return true;
            board[row+1][col] = temp;
            //System.out.println("XX");
        }
        if(col-1>=0&&board[row][col-1]!='0'&&t.next[board[row][col-1]-'a']!=null){
            char temp = board[row][col-1];
            Trie wh = t.next[board[row][col-1]-'a'];
            board[row][col-1] = '0';
            dfs(board, row, col-1, wh);
            //if(dfs(board, row-1, col, t))   return true;
            board[row][col-1] = temp;
            //System.out.println("XXX");
        }
        if(col+1<n&&board[row][col+1]!='0'&&t.next[board[row][col+1]-'a']!=null){
            char temp = board[row][col+1];
            Trie wh = t.next[board[row][col+1]-'a'];
            board[row][col+1] = '0';
            dfs(board, row, col+1, wh);
            //if(dfs(board, row-1, col, t))   return true;
            board[row][col+1] = temp;
            //System.out.println("XXXX");
        }
        //return false;
    }
}