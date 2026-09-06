namespace Solution.Tests;

public class SolutionTests
{
    [Theory]
    [InlineData(new[] { 3, 9, 7 }, 5, 4)]
    [InlineData(new[] { 4, 1, 3 }, 4, 0)]
    [InlineData(new[] { 3, 2 }, 6, 5)]
    public void MinOperations_ReturnsExpectedResult(int[] nums, int k, int expected)
    {
        var solution = new Solution();

        var actual = solution.MinOperations(nums, k);

        Assert.Equal(expected, actual);
    }
}
