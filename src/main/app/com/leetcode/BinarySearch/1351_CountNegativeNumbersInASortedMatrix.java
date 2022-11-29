package com.leetcode.BinarySearch;

class Solution62 {
    public int countNegatives(int[][] grid) {
        int count = 0;
        for(int i=0; i<grid.length; i++) {
            int l=0, h = grid[i].length-1, mid =0;
            boolean flag = false;
            while(l <= h) {
                mid = (h-l)/2+l;
                if(grid[i][mid] < 0) {
                    h = mid-1;
                    flag = true;
                } else {
                    l = mid+1;
                }
            }
            if(flag == true) {
                count += grid[i].length-l;
            }
        }
        return count;
    }
}