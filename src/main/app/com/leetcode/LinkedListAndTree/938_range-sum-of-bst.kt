package com.leetcode.LinkedListAndTree

import com.leetcode.TreeNode

class `938_range-sum-of-bst` {
    fun rangeSumBST(root: TreeNode?, low: Int, high: Int): Int {
        if (root == null) return 0
        var sum = 0

        print(root.`val`)
        if (root.`val` in low..high) {
            sum += root.`val`
            sum += rangeSumBST(root.left, low, high)
            sum += rangeSumBST(root.right, low, high)
        } else if (root.`val` > high) {
            sum += rangeSumBST(root.left, low, high)
        } else if (root.`val` < low) {
            sum += rangeSumBST(root.right, low, high)
        }

        return sum
    }
}