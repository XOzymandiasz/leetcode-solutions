namespace Solution.Tests;

public class SolutionTests
{
    public static TheoryData<int, IList<IList<int>>> TestCases => new()
    {
        {
            1,
            new List<IList<int>>
            {
                new List<int> { 1 }
            }
        },
        {
            2,
            new List<IList<int>>
            {
                new List<int> { 1 },
                new List<int> { 1, 1 }
            }
        },
        {
            3,
            new List<IList<int>>
            {
                new List<int> { 1 },
                new List<int> { 1, 1 },
                new List<int> { 1, 2, 1 }
            }
        },
        {
            5,
            new List<IList<int>>
            {
                new List<int> { 1 },
                new List<int> { 1, 1 },
                new List<int> { 1, 2, 1 },
                new List<int> { 1, 3, 3, 1 },
                new List<int> { 1, 4, 6, 4, 1 }
            }
        }
        
    };

    [Theory]
    [MemberData(nameof(TestCases))]
    public void Solve_ReturnsExpectedResult(int param, IList<IList<int>> expected)
    {
        var solution = new Solution();

        var actual = solution.Generate(param);

        Assert.Equal(expected.Count, actual.Count);

        for (var i = 0; i < expected.Count; i++)
        {
            Assert.Equal(expected[i], actual[i]);
        }
    }
}