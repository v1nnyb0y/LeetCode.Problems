package com.leetcode.Strings;

class Solution_1061 {
    private class UnionFind {
        private int[] parentArr;
        private UnionFind(int n) {
            parentArr = new int[n];
            for (int i = 0; i < n; i++) {
                parentArr[i] = i;
            }
        }

        private int getParent(int i) {
            return (parentArr[i] == i)? i : (parentArr[i] = getParent(parentArr[i]));
        }

        private void union(int i, int j) {
            int parent1 = getParent(i);
            int parent2 = getParent(j);
            if (parent1 < parent2) {
                parentArr[parent2] = parent1;
            } else {
                parentArr[parent1] = parent2;
            }
        }
    }

    public String smallestEquivalentString(String s1, String s2, String baseStr) {
        UnionFind uf = new UnionFind(26);
        for (int i = 0; i < s1.length(); i++) {
            int c1 = s1.charAt(i) - 'a';
            int c2 = s2.charAt(i) - 'a';
            uf.union(c1, c2);
        }

        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < baseStr.length(); i++) {
            int smallestMappedChar = uf.getParent(baseStr.charAt(i) - 'a');
            sb.append((char) (smallestMappedChar + 'a'));
        }

        return sb.toString();
    }
}