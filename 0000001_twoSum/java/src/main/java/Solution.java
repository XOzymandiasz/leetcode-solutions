import java.util.HashMap;
import java.util.Map;

class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<>();

        for (int i = 0; i < nums.length; i++) {
            Integer complement = target - nums[i];
            Integer previous = map.get(complement);

            if (previous != null) {
                return new int[]{previous, i};
            }

            map.put(nums[i], i);
        }

        return new int[0];
    }
}