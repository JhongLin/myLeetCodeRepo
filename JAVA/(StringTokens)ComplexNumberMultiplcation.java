class Solution {
    public String complexNumberMultiply(String num1, String num2) {
        String[] parts1 = num1.split("\\+");
        String[] parts2 = num2.split("\\+");
        /*
        for(int i = 0 ; i<parts1.length ; i++){
            System.out.println(parts1[i]);
        }*/
        int[] n1 = new int[2];
        int[] n2 = new int[2];
        n1[0] = str2int(parts1[0]);
        n1[1] = cpx2int(parts1[1]);
        n2[0] = str2int(parts2[0]);
        n2[1] = cpx2int(parts2[1]);
        //System.out.println(Arrays.toString(n1));
        //System.out.println(Arrays.toString(n2));
        int[] ans = new int[2];
        ans[0] = n1[0]*n2[0]-(n1[1]*n2[1]);
        ans[1] = n1[1]*n2[0]+n2[1]*n1[0];
        return new String(Integer.toString(ans[0])+'+'+Integer.toString(ans[1])+'i');
    }
    private int str2int(String a){
        if(a.charAt(0)=='-'){
            int count = 1, val = 0;
            for(int i = a.length()-1 ; i>0 ; i--){
                val+=(a.charAt(i)-'0')*count;
                count*=10;
            }
            return val*(-1);
        }else{
            int count = 1, val = 0;
            for(int i = a.length()-1 ; i>=0 ; i--){
                val+=(a.charAt(i)-'0')*count;
                count*=10;
            }
            return val;
        }
    }
    private int cpx2int(String b){
        if(b.charAt(0)=='-'){
            int count = 1, val = 0;
            for(int i = b.length()-2 ; i>0; i--){
                val+=(b.charAt(i)-'0')*count;
                count*=10;
            }
            return val*(-1);
        }else{
            int count = 1, val = 0;
            for(int i = b.length()-2 ; i>=0; i--){
                val+=(b.charAt(i)-'0')*count;
                count*=10;
            }
            return val;
        }
    }
}