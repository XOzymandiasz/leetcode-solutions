namespace Solution;

public class Solution
{
    public IList<int> GetRow(int rowIndex)
    {
        var row = new List<int>();
        
        var maxRowIndex = rowIndex + 1;
        long currentNumber = 1;
        row.Add((int)currentNumber);
        for (var i = 1; i < maxRowIndex; i++)
        {
            currentNumber *= maxRowIndex - i;
            currentNumber /= i;
            row.Add((int)currentNumber);
        }
        
        return row;
    }
}
