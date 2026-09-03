namespace Solution.Tests;

public class SolutionTests
{
    public static TheoryData<int, IList<int>> TestCases => new()
    {
        {
            0,
            new List<int>
            {
                1
            }
        },
        {
            1,
            new List<int>
            {
                1, 1
            }
        },
        {
            2,
            new List<int>
            {
                1, 2, 1
            }
        },
        {
            3,
            new List<int>
            {
               1, 3, 3, 1
            }
        },
        {
            4,
            new List<int>
            {
                1, 4, 6, 4, 1 
            }
        },
        {
            30,
            new List<int>
            {
                1,30,435,4060,27405,142506,593775,2035800,5852925,14307150,30045015,54627300,86493225,119759850,145422675,155117520,145422675,119759850,86493225,54627300,30045015,14307150,5852925,2035800,593775,142506,27405,4060,435,30,1
            }
        }
    };

    [Theory]
    [MemberData(nameof(TestCases))]
    public void Solve_ReturnsExpectedResult(int param, IList<int> expected)
    {
        var solution = new Solution();

        var actual = solution.GetRow(param);

        Assert.Equal(expected.Count, actual.Count);

        for (var i = 0; i < expected.Count; i++)
        {
            Assert.Equal(expected[i], actual[i]);
        }
    }
}