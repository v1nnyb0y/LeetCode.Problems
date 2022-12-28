package com.leetcode.LinkedListAndTree;

import com.leetcode.TreeNode;

class Solution134 {
    public TreeNode invertTree(TreeNode root) {
        if(root == null){
            return root;
        }

        //swap root's left and right node
        TreeNode temp = root.left;
        root.left = root.right;
        root.right = temp;

        //call left ans right nodes of tree
        invertTree(root.right);
        invertTree(root.left);

        return root;
    }
}