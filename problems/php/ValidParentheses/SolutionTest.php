<?php
namespace Problem\ValidParentheses;

use PHPUnit\Framework\TestCase;

class SolutionTest extends TestCase {
    public function testIsValid(): void {
        $solution = new Solution();

        $this->assertTrue($solution->isValid("()"));
        $this->assertTrue($solution->isValid("()[]{}"));
        $this->assertFalse($solution->isValid("(]"));
        $this->assertTrue($solution->isValid("()"));
    }
}