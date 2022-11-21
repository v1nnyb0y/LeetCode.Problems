class Solution {
    public String freqAlphabets(String s) {
        StringBuilder sb = new StringBuilder();

        for (int i = s.length() - 1; i >= 0; i--) {
            if (s.charAt(i) == '#') {
                int num = s.charAt(--i) - '0' + (s.charAt(--i) - '0') * 10;
                sb.append((char)(num + 96));
            } else {
                sb.append((char)(s.charAt(i) + 48));
            }
        }

        return sb.reverse().toString();
    }
}