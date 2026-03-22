import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.MethodSource;

import java.util.stream.Stream;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    static Stream<TestCase> cases() {
        return Stream.of(
                new TestCase("basic", new int[]{2, 7, 11, 15}, 9),
                new TestCase("unordered", new int[]{3, 2, 4}, 6),
                new TestCase("duplicates", new int[]{3, 3}, 6),
                new TestCase("negative", new int[]{7, 3, 4, -1}, 2)
        );
    }

    @ParameterizedTest(name = "{0}")
    @MethodSource("cases")
    void testTwoSum(TestCase testCase) {
        Solution solution = new Solution();

        int[] result = solution.twoSum(testCase.nums, testCase.target);

        assertEquals(2, result.length);
        assertNotEquals(result[0], result[1]);
        assertEquals(
                testCase.nums[result[0]] + testCase.nums[result[1]],
                testCase.target
        );
    }

    @org.junit.jupiter.api.Test
    void returnsEmptyArrayWhenNoSolutionExists() {
        Solution solution = new Solution();

        int[] result = solution.twoSum(new int[]{1, 2, 3}, 100);

        assertEquals(0, result.length);
    }

    static class TestCase {
        String name;
        int[] nums;
        int target;

        TestCase(String name, int[] nums, int target) {
            this.name = name;
            this.nums = nums;
            this.target = target;
        }

        @Override
        public String toString() {
            return name;
        }
    }
}