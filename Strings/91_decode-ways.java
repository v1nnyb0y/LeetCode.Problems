class Solution {
    public int numDecodings(String s) {
        if (s.charAt(0) == '0')
            return 0;
        var oneBack = 1;

        for (int i = 1, twoBack = 1; i < s.length(); i++) {
            var current = (s.charAt(i) == '0') ? 0 : oneBack;
            var pair = Integer.parseInt(s.substring(i - 1, i + 1));

            if (10 <= pair && pair <= 26)
                current += twoBack;

            twoBack = oneBack;
            oneBack = current;
        }
        return oneBack;
    }
}