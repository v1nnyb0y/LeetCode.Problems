class Solution {
    public int shortestPathLength(int[][] graph) {
        int n = graph.length;
        int mask = (1<<n) - 1;
        boolean[][] visited = new boolean[n][1<<n];
        Queue<Integer> queue = new LinkedList<>();
        for (int i = 0; i < n; i++) {
            int state = (i<<n) | (1<<i);
            queue.add(state);
            visited[i][1<<i] = true;
        }
        int dep = 0;
        while (queue.size() > 0) {
            int size = queue.size();
            while (size-- > 0) {
                int state = queue.poll();
                int node = state >> n;
                int status = state & mask;
                if (status == mask) return dep;
                for (int next: graph[node]) {
                    int newStatus = (1<<next) | status;
                    int newState = (next<<n) | newStatus;
                    if (!visited[next][newStatus]) {
                        queue.add(newState);
                        visited[next][newStatus] = true;
                    }
                }
            }
            dep++;
        }
        return -1;
    }
}