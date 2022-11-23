import java.lang.Math;

class Solution {
    int maxd = 0;
    public void visit(TreeNode n, int depth) {
        if (n == null)
            return;

        maxd = Math.max(maxd, depth + 1);
        visit(n.left, depth + 1);
        visit(n.right, depth + 1);
    }

    public int maxDepth(TreeNode root) {
        visit(root, 0);
        return maxd;
    }
}