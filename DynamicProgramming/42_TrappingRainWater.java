class Solution {
    public int trap(int[] height) {
        int n = height.length;
        if(n <=2) {return 0;}

        int[] left_max = new int[n];
        int[] right_max = new int[n];
        int[] store = new int[n];

        int left_temp = 0;
        int right_temp = 0;

        for(int i = 0; i < n; i++){
            left_temp = Math.max(left_temp, height[i]);
            left_max[i] = left_temp;
        }
        for(int i = n-1; i>=0; i--){
            right_temp = Math.max(right_temp, height[i]);
            right_max[i] = right_temp;
        }

        store[0] = 0;
        store[n - 1] = 0;
        for(int i = 1; i<= n-2; i++){
            int low = Math.min(left_max[i-1], right_max[i+1]);
            store[i] = (low > height[i])? (low - height[i]) : 0 ;
        }

        int sum = 0;
        for(int s : store){
            sum = sum + s;
        }
        return sum;
    }
}