package com.leetcode.Array;

class Solution_134 {
    public int canCompleteCircuit(int[] gas, int[] cost) {
        int start = 0;
        int curTank = 0, tank = 0;
        for(int i = 0; i < gas.length; i++){
            int curGain = gas[i] - cost[i];
            tank += curGain;
            if(curTank + curGain < 0){
                start = i + 1;
                curTank = 0;
            }
            else
                curTank += curGain;
        }
        if(tank < 0)
            return -1;
        else
            return start;
    }
}
