class MedianFinder {
    Queue<Integer> small ;
    Queue<Integer> big ;

    public MedianFinder() {
        small = new PriorityQueue<Integer>((n1 , n2) -> n2-n1);
        big = new PriorityQueue<Integer>((n1 , n2) -> n1-n2);
    }

    public void addNum(int num) {
        small.add(num);

        if(!small.isEmpty() && !big.isEmpty() && small.peek() > big.peek())
            big.add(small.remove());

        if(small.size() > big.size()+1)
            big.add(small.remove());

        if(big.size() > small.size()+1)
            small.add(big.remove());
    }

    public double findMedian() {
        if(small.size() > big.size())
            return small.peek();
        if(big.size() > small.size())
            return big.peek();

        return (small.peek() + big.peek()) / 2d;
    }
}