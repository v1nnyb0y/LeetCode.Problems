package com.leetcode.TwoPointers;

class ListNode82 {
    ListNode82 next;
    int val;

    public ListNode82(int value) {
        val = value;
    }
}

class Solution_82 {
    public ListNode82 deleteDuplicates(ListNode82 head) {
        if(head == null || head.next == null) return head;
        ListNode82 dummy = new ListNode82(0), l = dummy;
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