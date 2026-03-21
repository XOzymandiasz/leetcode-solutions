<?php

class Solution {
    /**
     * @param Integer[] $nums
     * @return Integer[]
     */
    function twoSum(array $nums, int $target): array {
        $visited = [];
        foreach ($nums as $i => $num) {
            if (isset($visited[$target - $num])) {
                return [$visited[$target - $num], $i];
            }
            $visited[$num] = $i;
        }
        return [];
    }
}
