class Solution {
    public int minimumTotal(List<List<Integer>> triangle) {
        for(int i=1;i<triangle.size();i++){
            for(int j=0;j<=i;j++){
                triangle.get(i).set(j,Math.min((j==i?100000:triangle.get(i-1).get(j)),(j==0?100000:triangle.get(i-1).get(j-1)))+triangle.get(i).get(j));
            }
        }
        int min=100000;
        for(int i=0;i<triangle.size();i++) min = Math.min(min,triangle.get(triangle.size()-1).get(i));
        return min;
    }
}