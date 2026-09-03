namespace Solution;

public class Solution
{
    public IList<IList<int>> Generate(int numRows)
    {
        var triangle = new List<IList<int>>
        {
            new List<int> { 1 }
        };
        
        for (var i = 1; i < numRows; i++)
        {
            var row = new List<int>();
            var previousRow = triangle[i-1];
            row.Add(1);
            for (var j = 0; j < i - 1; j++)
            {
                var value = previousRow[j] + previousRow[j + 1];
                row.Add(value);
            }
            row.Add(1);
            triangle.Add(row);
        }
        
        return triangle;
    }
}
