// https://leetcode.com/problems/prefix-and-suffix-search/



typedef struct trie{
    struct trie *ch[27];
    int index;
}Trie;


typedef struct {
    Trie *root;
} WordFilter;


void insert(Trie *node, char *str, int idx)
{
    //printf("Insert: %s\n", str);
    for(int i = 0 ; i<strlen(str) ; i++)
    {
        if(!node->ch[ str[i]-'a' ])
            node->ch[str[i]-'a'] = (Trie*)calloc(1, sizeof(Trie));
        node = node->ch[str[i]-'a'];
        node->index = idx;
    }
}

WordFilter* wordFilterCreate(char ** words, int wordsSize) {
    //Initialization
    WordFilter *new = (WordFilter*)malloc(sizeof(WordFilter));
    Trie *newT = (Trie*)calloc(1, sizeof(Trie) );
    new->root = newT;
    
    //Input the words to dictionary
    for(int i = 0 ; i<wordsSize ; i++)
    {
        int Len = strlen(words[i]);
        insert(new->root, words[i], i);
        char *temp = (char*)calloc(Len*2+2, sizeof(char));
        strncpy(temp, words[i], Len);
        temp[Len] = '{';
        strncpy(&temp[Len+1], words[i], Len);
        //printf("%s\n", temp);
        for(int j = 1 ; j<=Len ;j++)
        {
            insert(new->root, &temp[Len-j], i);
            //printf("%s\n", &temp[Len-j]);
        }
        
        free(temp);
    }
    
    
    return new;
}

int wordFilterF(WordFilter* obj, char * prefix, char * suffix) {
    
    char *temp = (char*)calloc(strlen(prefix)+strlen(suffix)+2, sizeof(char));
    strncpy(temp, suffix, strlen(suffix));
    
    temp[ strlen(suffix) ] = '{';
    strncpy(&temp[strlen(suffix)+1], prefix, strlen(prefix));
    
    //printf("Search: %s\n", temp);
    Trie *node = obj->root;
    for(int i = 0 ; i<strlen(temp) ; i++)
    {
        if(!node->ch[temp[i]-'a'])  return -1;
        node = node->ch[temp[i]-'a'];
    }
    return node->index;
    
    return 0;
}

void freeTrie(Trie *t)
{
    for(int i = 0 ; i<27 ; i++ )
    {
        if(t->ch[i])    freeTrie(t->ch[i]);
    }
    free(t);
}

void wordFilterFree(WordFilter* obj) {
    freeTrie(obj->root);
    free(obj);
}

/**
 * Your WordFilter struct will be instantiated and called as such:
 * WordFilter* obj = wordFilterCreate(words, wordsSize);
 * int param_1 = wordFilterF(obj, prefix, suffix);
 
 * wordFilterFree(obj);
*/