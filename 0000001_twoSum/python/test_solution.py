import pytest
from solution import Solution

cases = [
    ("basic", [2, 7, 11, 15], 9),
    ("unordered", [3, 2, 4], 6),
    ("duplicates", [3, 3], 6),
    ("negative", [7, 3, 4, -1], 2),
]

@pytest.mark.parametrize("name, nums, target", cases)
def test_two_sum(name, nums, target):
    sol = Solution()
    i, j = sol.twoSum(nums, target)

    assert i != j
    assert nums[i] + nums[j] == target


def test_no_solution():
    sol = Solution()
    result = sol.twoSum([1, 2, 3], 100)

    assert result == []