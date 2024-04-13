package com.leetcode.LinkedListAndTree;

import com.leetcode.TreeNode;

class Solution_124 {
    int max=Integer.MIN_VALUE;
    public int utilfunc(TreeNode root){
        if(root==null)
            return 0;
        int l=utilfunc(root.left);
        int r=utilfunc(root.right);
        max=Math.max(max,Math.max(Math.max(l,r),Math.max(l+r,0))+root.val);
        return Math.max(l,Math.max(r,0))+root.val;
    }
    public int maxPathSum(TreeNode root) {
        utilfunc(root);
        return max;
    }
}