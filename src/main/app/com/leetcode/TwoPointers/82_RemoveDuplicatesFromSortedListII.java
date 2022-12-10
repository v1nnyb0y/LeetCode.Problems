package com.leetcode.TwoPointers;

class Solution_82 {
    public ListNode deleteDuplicates(ListNode head) {
        if(head == null || head.next == null) return head;
        ListNode dummy = new ListNode(0), l = dummy;
        dummy.next = head;
        while(head != null && head.next != null){
            if(head != null && head.next != null && head.val != head.next.val){
                l = l.next;
            }
            while(head != null && head.next != null && head.val == head.next.val){
                l.next = head.next.next;
                head = head.next;
            }
            head = head.next;
        }
        return dummy.next;
    }
}