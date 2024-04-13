package com.leetcode.Strings;

class Solution11 {
    public String interpret(String command) {
        if(command.isBlank() || command.length()==1){
            return command;
        }

        StringBuilder sb = new StringBuilder();
        int length = command.length();
        for(int i=0;i<length;i++){
            char c = command.charAt(i);
            if(c=='G'){
                sb.append('G');
            }else if(c=='(' && i<length-1&& command.charAt(i+1)==')'){
                sb.append("o");
                i++;
            }else{
                sb.append("al");
                i+=3;
            }
        }
        return sb.toString();
    }
}