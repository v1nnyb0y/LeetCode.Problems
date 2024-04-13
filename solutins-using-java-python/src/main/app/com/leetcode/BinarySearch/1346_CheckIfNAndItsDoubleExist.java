package com.leetcode.BinarySearch;

import java.util.HashSet;
import java.util.Set;

class Solution55 {
    public boolean checkIfExist(int[] arr) {
        Set<Integer> set=new HashSet<>();
        for(int num : arr){
            if(num%2==0 && set.contains(num/2))return true;
            if(set.contains(num*2))return true;
            set.add(num);
        }
        return false;
    }
}
