package com.leetcode;

import java.util.PriorityQueue;

public class Main {
    public static void main(String[] args) {
        PriorityQueue<Integer> pQueue = new PriorityQueue<>();
        pQueue.add(1);
        pQueue.add(5);
        pQueue.add(2);
        pQueue.add(4);
        pQueue.add(3);

        int counter = 3;
        while (counter > 0) {
            pQueue.add(counter * 12);
            System.out.println(pQueue.poll());
            counter -= 1;
        }

        while (!pQueue.isEmpty()) {
            System.out.println(pQueue.poll());
        }
    }
}