// twoSum returns indices of two numbers whose sum equals the target.
// Uses a map to store visited numbers and find complements in O(1),
// resulting in an overall O(n) time complexity and O(n) space complexity.
export function twoSum(nums: number[], target: number): number[] {
    const visited: Record<number, number> = {};

    for (let i = 0; i < nums.length; i++) {
        const complement = target - nums[i];
        if (complement in visited) {
            return [visited[complement], i];
        }
        visited[nums[i]] = i;
    }

    return []
}