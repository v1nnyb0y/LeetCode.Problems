package com.leetcode.LinkedListAndTree;

import java.util.HashSet;
import java.util.Set;

/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
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