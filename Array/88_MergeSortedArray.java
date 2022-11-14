class Solution {
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        int i = m-1;
        int j = n-1;
        int idx = nums1.length-1;
        //start looping from last values of both arrays
        //if elements remains in both, compare them and put larger
        while(i>=0 && j>=0){
            if(nums1[i]> nums2[j]){
                nums1[idx] = nums1[i];
                i--;
            }
            else{
                nums1[idx] = nums2[j];
                j--;
            }
            idx--;
        }
        //if elements are remaining in one array, fill them
        while(i>=0){
            nums1[idx] = nums1[i];
            i--;
            idx--;
        }
        while(j>=0){
            nums1[idx] = nums2[j];
            j--;
            idx--;
        }
    }
}