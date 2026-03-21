/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
export var twoSum = function(nums, target) {
    const visited = new Map();

    for (let i = 0; i < nums.length; i++) {
        const complement = target - nums[i];

        if (visited.has(complement)) {
            return [visited.get(complement), i];
        }

        visited.set(nums[i], i);
    }

    return [];
};