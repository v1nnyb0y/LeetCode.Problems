class Solution {
    public List<List<Integer>> levelOrder(TreeNode root) {
        List<List<Integer>> res = new LinkedList<List<Integer>>();
        List<Integer> row = new LinkedList<>();
        Queue<TreeNode> q = new LinkedList<>();
        TreeNode t=null;
        if(root==null)
            return res;
        q.add(root);
        row.add(root.val);
        res.add(row);
        row = new LinkedList<>();
        q.add(null);

        while(!q.isEmpty()){
            while((t=q.poll())!=null){

                if(t.left != null){
                    q.add(t.left);
                    row.add(t.left.val);
                }

                if(t.right!=null){
                    q.add(t.right);
                    row.add(t.right.val);
                }
            }
            if(!q.isEmpty()) {
                q.add(null);
                res.add(row);
            }
            row = new LinkedList<Integer>();
        }
        return res;
    }
}