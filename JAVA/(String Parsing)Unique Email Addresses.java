//https://leetcode.com/problems/unique-email-addresses/

class Solution {
    public int numUniqueEmails(String[] emails) {
        HashSet<String> repository = new HashSet<>();
        for(int i = 0 ; i < emails.length ; i++){
            char[] handling = emails[i].toCharArray();
            List<Character> build = new ArrayList<>();
            int j = 0;
            while(handling[j]!='@'){
                if(handling[j]=='+'){
                    while(handling[j]!='@') j++;
                    break;
                }else if(handling[j]!='.'){
                    build.add(handling[j]);
                }
                j++;
            }
            build.add(handling[j]);
            j++;
            while(j<handling.length){
                if(handling[j]=='+'){
                    break;
                }else{
                    build.add(handling[j]);
                }
                j++;
            }
            StringBuilder myBuilder = new StringBuilder(build.size());
            for(Character ch : build)    myBuilder.append(ch);
            String str = myBuilder.toString();
            if(!repository.contains(str))   repository.add(str);
        }
        //System.out.println(repository.toString());
        return repository.size();
    }
}