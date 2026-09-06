namespace Solution;

public class Solution
{
    public int[] FindDegrees(int[][] matrix) 
    {
        return matrix
            .Select(row => row.Sum())
            .ToArray();
    }
}
