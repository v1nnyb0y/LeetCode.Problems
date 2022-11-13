class Solution {
    public int tribonacci(int n) {
        int[] tribSec = new int[n + 3];
        tribSec[0] = 0;
        tribSec[1] = 1;
        tribSec[2] = 1;

        for(var i = 3; i < tribSec.length; i++) {
            tribSec[i] = tribSec[i - 3] + tribSec[i - 2] + tribSec[i - 1];
        }

        return tribSec[n];
    }
}