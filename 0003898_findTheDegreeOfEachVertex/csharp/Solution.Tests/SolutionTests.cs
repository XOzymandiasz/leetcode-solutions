namespace Solution.Tests;

public class SolutionTests
{
    public static TheoryData<int[][], int[]> TestCases => new()
    {
        {
            [
                [ 0, 1, 1 ],
                [1, 0, 1 ],
                [1, 1, 0]
            ],
            [2, 2, 2]
        },
        {
            [
                [0, 0, 0],
                [0, 0, 0],
                [0, 0, 0]
            ],
            [0, 0, 0]
        },
        {
            [
                [0, 1, 1, 1],
                [1, 0, 0, 0],
                [1, 0, 0, 0],
                [1, 0, 0, 0]
            ],
            [3, 1, 1, 1]
        }
    };

    [Theory]
    [MemberData(nameof(TestCases))]
    public void FindDegrees_ReturnsExpectedResult(
        int[][] matrix,
        int[] expected)
    {
        var solution = new Solution();

        var actual = solution.FindDegrees(matrix);

        Assert.Equal(expected, actual);
    }
}
