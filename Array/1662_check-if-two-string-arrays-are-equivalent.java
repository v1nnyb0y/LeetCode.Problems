class Solution {
  public boolean arrayStringsAreEqual(String[] word1, String[] word2) {
    int a = 0, b = 0, x = 0, y = 0;

    while(a != word1.length && b != word2.length){
      if(word1[a].charAt(x) != word2[b].charAt(y)) return false;

      if(++x == word1[a].length()) {x = 0; a++;}
      if(++y == word2[b].length()) {y = 0; b++;}
    }

    return a == word1.length && b == word2.length;
  }
}