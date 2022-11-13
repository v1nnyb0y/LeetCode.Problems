class Solution {
    public double average(int[] salaries) {
        int max = Integer.MIN_VALUE, min = Integer.MAX_VALUE, sum = 0, count = salaries.length - 2;
        for(int salary: salaries) {
            if (salary > max) max = salary;
            if (salary < min) min = salary;
            sum += salary;
        }

        return (double) (sum - max - min) / count;
    }
}