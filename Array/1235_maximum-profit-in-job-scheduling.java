class Solution {

    //make a class so that, everything of a job can be access at once
    public class job{
        int start;
        int end;
        int profit;
        public job(int start,int end,int profit){
            this.start = start;
            this.end = end;
            this.profit = profit;
        }
    }

    public int jobScheduling(int[] startTime, int[] endTime, int[] profit) {
        job[] arr = new job[startTime.length];
        for(int i=0;i<arr.length;i++){
            arr[i] = new job(startTime[i],endTime[i],profit[i]);
        }
        //sort the jobs on end time basis
        Arrays.sort(arr,(a,b)->Integer.compare(a.end,b.end));

        return func(arr);
    }


    public int func(job[] arr){
        //dp will contain best project till each job
        int dp[] = new int[arr.length];

        dp[0] = arr[0].profit;
        //omax will be our global max
        int omax = 0;
        //for each job
        for(int i=1; i < arr.length; i++){
            dp[i] = Math.max(arr[i].profit, dp[i-1]);
            //check its previous jobs
            for(int j=i-1; j>=0 ; j--){
                //if they dont over lap, take max
                if(arr[j].end <= arr[i].start){
                    dp[i] = Math.max(dp[i], arr[i].profit + dp[j]);
                    break;
                }
            }
            omax = Math.max(omax,dp[i]);
        }
        return omax;
    }

}