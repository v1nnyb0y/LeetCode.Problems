import java.util.ArrayList;
import java.util.List;

class Solution {
    public List<Integer> getRow(int rowIndex) {

        List<Integer> al = new ArrayList<>();
        al.add(1);
        for (int i = 1; i <= rowIndex; i++) {
            ArrayList<Integer> al2 = new ArrayList<>();
            al2.add(1);
            for (int j = 0; j < al.size() - 1; j++)
                al2.add(al.get(j) + al.get(j + 1));
            al2.add(1);
            al = al2;
        }
        return al;
    }
}