#include <iostream>
#include <vector>
#include <unordered_map>

#include "solution.hpp"

// twoSum returns indices of two numbers whose sum equals the target.
// Uses a hash map to store visited numbers and find complements in O(1),
// resulting in an overall O(n) time complexity and O(n) space complexity.
std::vector<int> Solution::twoSum(std::vector<int>& nums, int target) {
    std::unordered_map<int, int> m;

    for (size_t i = 0; i < nums.size(); ++i) {
        int complement = target - nums[i];

        if (m.find(complement) != m.end()) {
            return {m[complement], static_cast<int>(i)};
        }

        m[nums[i]] = i;
    }

    return {};
}