package com.leetcode.Strings;

import java.util.ArrayList;
import java.util.List;

class Solution3 {
    public List<Integer> findAnagrams(String s, String p) {
        List<Integer> res = new ArrayList<>();
        if(p.length() > s.length()) return res;

        int[] arr = getCharArray(s, p);
        if(checkAnagram(arr)) res.add(0);

        int start = 0;
        int end = p.length()-1;

        while(end < s.length()-1){
            // remove start element
            arr[s.charAt(start) - 'a']--;
            start++;
            // add end+1 element
            end++;
            arr[s.charAt(end) - 'a']++;

            if(checkAnagram(arr)) res.add(start);
        }

        return res;
    }

    boolean checkAnagram(int[] arr){
        for(int i=0; i<arr.length; i++){
            if(arr[i] != 0){
                return false;
            }
        }
        return true;
    }

    int[] getCharArray(String s, String p){
        int[] arr = new int[26];
        for(int i=0; i<p.length(); i++){
            arr[s.charAt(i) - 'a']++;
            arr[p.charAt(i) - 'a']--;
        }
        return arr;
    }
}