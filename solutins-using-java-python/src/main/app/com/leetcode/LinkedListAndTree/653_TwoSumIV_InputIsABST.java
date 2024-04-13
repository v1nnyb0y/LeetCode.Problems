package com.leetcode.LinkedListAndTree;

import com.leetcode.TreeNode;

import java.util.HashSet;
import java.util.Set;

class Solution114 {
    public boolean findTarget(TreeNode root, int k) {
        return findTarget(root, k, new HashSet<>());
    }

    private boolean findTarget(TreeNode root, int k, Set<Integer> seen) {
        if (root == null)
            return false;
        if (seen.contains(k - root.val))
            return true;

        seen.add(root.val);

        return findTarget(root.left, k, seen)
                || findTarget(root.right, k, seen);
    }
}