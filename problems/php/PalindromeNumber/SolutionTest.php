<?php
namespace Problem\PalindromeNumber;

use PHPUnit\Framework\TestCase;

class SolutionTest extends TestCase {
    public function testIsPalindrome(): void {
        $solution = new Solution();

        $this->assertTrue($solution->isPalindrome(121));
        $this->assertFalse($solution->isPalindrome(-121));
        $this->assertFalse($solution->isPalindrome(10));
    }
}