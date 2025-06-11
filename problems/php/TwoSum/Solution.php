<?php

namespace Problem\TwoSum;

class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer[]
     */
    function twoSum($nums, $target) {
        $diffs = [];

        foreach ($nums as $numKey => $num) {
            $diff = $target - $num;
            if (isset($diffs[$num])) {
                return [$diffs[$num], $numKey];
            }

            $diffs[$diff] = $numKey;
        }

        return [];
    }
}