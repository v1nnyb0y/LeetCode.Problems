package com.leetcode.Strings;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;

class Solution_49 {
    public List<List<String>> groupAnagrams(String[] strs) {
        List<List<String>> list = new ArrayList<>();
        HashMap<String , ArrayList<String>> map = new HashMap<>();
        for(String s : strs){
            char[] ch = s.toCharArray();
            Arrays.sort(ch);
            String sorted_string = new String(ch);
            if(map.containsKey(sorted_string)){
                ArrayList<String> currlist = map.get(sorted_string);
                currlist.add(s);
                map.put(sorted_string,currlist);
            }
            else{
                ArrayList<String> li = new ArrayList();
                li.add(s);
                map.put(sorted_string,li);
            }
        }


        for(String s : map.keySet()){
            list.add(new ArrayList(map.get(s)));
        }

        return list;
    }
}