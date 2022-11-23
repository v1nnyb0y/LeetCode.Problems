class Solution {
    public int minReorder(int n, int[][] connections) {
        List<Integer> []direct = new ArrayList[n];
        List<Integer> []reverse = new ArrayList[n];
        for(int i=0;i<n;i++){
            direct[i] = new ArrayList();
            reverse[i] = new ArrayList();
        }
        for(int []a: connections){
            direct[a[0]].add(a[1]);
            reverse[a[1]].add(a[0]);
        }
        int ans = 0;
        Queue<Integer> q = new ArrayDeque();
        boolean []visited = new boolean[n];
        visited[0] = true;
        q.add(0);
        while(!q.isEmpty()){
            int i = q.poll();
            for(int j=0;j<direct[i].size();j++){
                if(!visited[direct[i].get(j)]){
                    ans ++;
                    visited[direct[i].get(j)] = true;
                    q.add(direct[i].get(j));
                }
            }
            for(int j=0;j<reverse[i].size();j++){
                if(!visited[reverse[i].get(j)]){
                    visited[reverse[i].get(j)] = true;
                    q.add(reverse[i].get(j));
                }
            }
        }
        return ans;
    }
}