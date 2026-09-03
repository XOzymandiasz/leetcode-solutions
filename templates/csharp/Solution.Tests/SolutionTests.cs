namespace Solution.Tests;

public class SolutionTests
{
    [Theory]
    [InlineData(1, 2)]
    public void Solve_ReturnsExpectedResult(int param, int expected)
    {
        var solution = new Solution();

        var actual = solution.Solve(param);

        Assert.Equal(expected, actual);
    }
}
