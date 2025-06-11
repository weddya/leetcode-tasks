<?php
namespace Problem\LongestCommonPrefix;

use PHPUnit\Framework\TestCase;

class SolutionTest extends TestCase {
    public function testLongestCommonPrefix(): void {
        $solution = new Solution();

        $this->assertEquals("fl", $solution->longestCommonPrefix(["flower", "flow", "flight"]));
        $this->assertEquals("", $solution->longestCommonPrefix(["dog","racecar","car"]));
    }
}