<?php
namespace Problem\PalindromeNumber;

class Solution {

    /**
     * @param Integer $x
     * @return Boolean
     */
    function isPalindrome(int $x): bool {
        if ($x < 0) {
            return false;
        }

        $originX = $x;
        $reversedX = 0;
        while ($x >= 1) {
            $revDigit = $x % 10;
            $reversedX = $reversedX * 10 + $revDigit;
            $x = ($x - $revDigit) / 10;
        }

        return $originX === $reversedX;
    }
}