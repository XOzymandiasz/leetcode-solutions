namespace DefaultNamespace;

public class Solution {
    // TwoSum returns indices of two numbers whose sum equals the target.
    // Uses a map to store visited numbers and find complements in O(1),
    // resulting in an overall O(n) time complexity and O(n) space complexity.
    public int[] TwoSum(int[] nums, int target)
    {
        Dictionary<int, int> map = new Dictionary<int, int>(nums.Length);
        
        map[nums[0]] = 0;  // it gives higher chance to Get in TryGetValue
        for (int i = 1; i < nums.Length; i++)
        {
            int complement = target - nums[i];

            if (map.TryGetValue(complement, out int previous))
            {
                return new int[2] { previous, i };
            }
            
            map[nums[i]] = i;
        }
        return Array.Empty<int>();
    }
}