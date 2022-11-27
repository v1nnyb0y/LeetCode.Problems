package com.leetcode.LinkedListAndTree;

class Node {
    Node left;
    Node right;
    Node next;
    int val;
}

class Solution101 {
    public Node connect(Node root) {
        if(root==null) return null;
        if(root.left!=null) root.left.next = root.right;
        if(root.right!=null) root.right.next = (root.next!=null)?root.next.left:null;
        connect(root.left);
        connect(root.right);
        return root;
    }
}