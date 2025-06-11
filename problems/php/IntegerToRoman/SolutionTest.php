<?php
namespace Problem\IntegerToRoman;

use PHPUnit\Framework\TestCase;

class SolutionTest extends TestCase {
    public function testIntToRoman(): void {
        $solution = new Solution();

        $this->assertEquals("MMMDCCXLIX", $solution->intToRoman(3749));
        $this->assertEquals("LVIII", $solution->intToRoman(58));
        $this->assertEquals("MCMXCIV", $solution->intToRoman(1994));
    }
}