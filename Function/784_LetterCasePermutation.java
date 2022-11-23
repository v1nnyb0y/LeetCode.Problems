import java.util.ArrayList;
import java.util.List;

public class LetterCasePermutation {
    List<String> results = new ArrayList<>();
    String s;

    public List<String> letterCasePermutation(String s) {
        this.s = s;
        backtrack(new StringBuilder(), 0);
        return results;

    }

    private void backtrack(StringBuilder sb, int index) {
        if (index == s.length()) {
            results.add(sb.toString());
            return;
        }
        char c = s.charAt(index);
        if (Character.isDigit(c)) {
            sb.append(c);
            backtrack(new StringBuilder(sb), index + 1);
        } else {
            sb.append(Character.toLowerCase(c));
            backtrack(new StringBuilder(sb), index + 1);
            sb.deleteCharAt(sb.length() - 1);
            sb.append(Character.toUpperCase(c));
            backtrack(new StringBuilder(sb), index + 1);
        }

    }
}