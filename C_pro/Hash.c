struct hash_table{
  int key;
  int index;
};


unsigned long hash(unsigned char *str)
{
    unsigned long hash = 5381;
    int c;

    while (c = *str++)
        hash = ((hash << 5) + hash) + c; /* hash * 33 + c */

    return hash;
}

int find_hash(struct hash_table *tab,int key,int size){

    for(int k=0;k<size;k++){
        if(tab[k].key==key)
           return k;
    }
   return -1;
}


// Complete the checkMagazine function below.
void checkMagazine(int magazine_count, char** magazine, int note_count, char** note, struct hash_table *tab) {

   for(int i=0;i<magazine_count;i++){
         int key=hash(magazine[i]);
         int k=find_hash(tab,key,magazine_count);
         if(k!=-1){
           tab[k].index++;
           tab[k].key=key;
         }
         else
         {
            tab[i].index=0;
            tab[i].index++;
            tab[i].key=key;
         }
   }

   for(int i=0;i<note_count;i++){
       unsigned long key=hash(note[i]);
       
       int k=find_hash(tab,key,magazine_count);
        if(k!=-1 && tab[k].index > 0)
             tab[k].index--;
                      
        else
        {  printf("No");
              return;
           
        }

   }
   printf("Yes");

}