class Solution {

    static int i=0;
    public ListNode removeNthFromEnd(ListNode head, int n) {
        //System.out.println(fun(head,n).val);
        return fun(head,n);
    }

    public ListNode fun(ListNode head,int x){
        if(head.next==null){
            i=1;
            if(i==x) return null;
            return head;
        }

        head.next=fun(head.next,x);
        i++;
        if(i==x)
            return head.next;
        else
            return head;
    }
}