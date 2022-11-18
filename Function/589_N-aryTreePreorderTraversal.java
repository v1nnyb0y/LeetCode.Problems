class Solution {
    public List<Integer> preorder(Node root) {
        List<Integer> res = new ArrayList<>();

        if(root == null){
            return res;
        }

        Stack<Node> stack = new Stack<>();
        stack.push(root);

        while(!stack.isEmpty()){
            Node current = stack.peek();
            stack.pop();

            res.add(current.val);

            for(int i = current.children.size()-1; i>=0; i--){
                stack.push(current.children.get(i));
            }
        }
        return res;
    }
}