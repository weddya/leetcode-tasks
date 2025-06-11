<?php

namespace Problem\TwoSum;

use PHPUnit\Framework\TestCase;

class SolutionTest extends TestCase {
    public function testTwoSum() {
        $solution = new Solution();

        $this->assertEquals([0, 1], $solution->twoSum([2, 7, 11, 15], 9));
        $this->assertEquals([1, 2], $solution->twoSum([3, 2, 4], 6));
        $this->assertEquals([0, 1], $solution->twoSum([3, 3], 6));
    }
}