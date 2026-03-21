#define CATCH_CONFIG_MAIN
#include <catch2/catch.hpp>
#include <vector>
#include <string>

#include "solution.hpp"

TEST_CASE("twoSum finds valid indices", "[twoSum]") {
    struct TestCase {
        std::string name;
        std::vector<int> nums;
        int target;
    };

    const std::vector<TestCase> cases = {
        {"basic",      {2, 7, 11, 15}, 9},
        {"unordered",  {3, 2, 4}, 6},
        {"duplicates", {3, 3}, 6},
        {"negative",   {7, 3, 4, -1}, 2},
    };

    Solution s;
    for (const auto& tc : cases) {
        SECTION(tc.name) {
            auto nums = tc.nums;
            auto result = s.twoSum(nums, tc.target);

            const int i = result[0];
            const int j = result[1];

            REQUIRE(i != j);
            REQUIRE(nums[i] + nums[j] == tc.target);
        }
    }
}

TEST_CASE("twoSum returns empty vector when no solution exists", "[twoSum]") {
    std::vector<int> nums = {1, 2, 3};
    
    Solution s;
    auto result = s.twoSum(nums, 100);

    REQUIRE(result.empty());
}