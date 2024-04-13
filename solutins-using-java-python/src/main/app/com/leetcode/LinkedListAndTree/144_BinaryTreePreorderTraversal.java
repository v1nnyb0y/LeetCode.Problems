package com.leetcode.LinkedListAndTree;

import com.leetcode.TreeNode;

import java.util.LinkedList;
import java.util.List;

class Solution102 {
    public List<Integer> preorderTraversal(TreeNode root) {
        List<Integer> ans = new LinkedList<>();
        LinkedList<TreeNode> st = new LinkedList<>();

        while(root != null || !st.isEmpty())
            if(root != null){
                ans.add(root.val);
                if(root.right != null) st.addLast(root.right);
                root = root.left;
            }
            else root = st.pollLast();

        return ans;
    }
}