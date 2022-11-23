import java.util.*;
class Solution {


    static  List<List<Integer>> list;
    static List<Integer> temp;
    public List<List<Integer>> combine(int n, int k) {
        list=new ArrayList<>();
        temp=new ArrayList<>();
        combination(1,n,k);
        return list;

    }
    void combination(int c,int n,int k){
        if(k==0){
            List<Integer> l=new ArrayList<>();
            for(int op:temp)
                l.add(op);
            list.add(l);
            return;
        }
        if(c>n)
            return;
        for(int i=c;i<=n;i++){
            temp.add(i);
            combination(i+1,n,k-1);
            temp.remove(temp.size()-1);
        }
    }

}