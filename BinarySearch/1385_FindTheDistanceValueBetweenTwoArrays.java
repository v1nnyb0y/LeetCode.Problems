class Solution {
    private int[] srtArr2;
    private boolean haveElems(int v,int d){
        if(srtArr2[0]>v+d||srtArr2[srtArr2.length-1]<v-d)
            return false;
        int p1=Arrays.binarySearch(srtArr2,v-d);
        int p2=Arrays.binarySearch(srtArr2,v+d);
        return p1>=0||p2>=0||p1!=p2;
    }
    public int findTheDistanceValue(int[] arr1, int[] arr2, int d) {
        Arrays.sort(arr2);
        srtArr2=arr2;
        return IntStream.of(arr1).filter(a->!haveElems(a,d)).map(a->a/a).sum();
    }
}