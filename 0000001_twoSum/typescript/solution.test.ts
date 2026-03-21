import { describe, it, expect } from 'vitest';
import { twoSum } from './solution';

describe('twoSum', () => {
    const cases = [
        { name: 'basic', nums: [2, 7, 11, 15], target: 9 },
        { name: 'unordered', nums: [3, 2, 4], target: 6 },
        { name: 'duplicates', nums: [3, 3], target: 6 },
        { name: 'negative', nums: [7, 3, 4, -1], target: 2 },
    ];

    cases.forEach(({ name, nums, target }) => {
        it(name, () => {
            const [i, j] = twoSum(nums, target);
            expect(i).not.toBe(j);
            expect(nums[i] + nums[j]).toBe(target);
        });
    });

    it('returns empty array when no solution exists', () => {
        const result = twoSum([1, 2, 3], 100);
        expect(result).toEqual([]);
    });
});