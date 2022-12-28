package com.leetcode.LinkedListAndTree;

import com.leetcode.TreeNode;

class Solution299 {
    public int countNodes(TreeNode root) {
        if (root == null) return 0;

        return countNodes(root.left) + countNodes(root.right) + 1;
    }
}