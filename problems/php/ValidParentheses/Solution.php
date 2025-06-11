<?php
namespace Problem\ValidParentheses;

class Solution {
    const CLOSING_BRACKETS = [
        ')' => '(',
        ']' => '[',
        '}' => '{',
    ];

    /**
     * @param String $s
     * @return Boolean
     */
    function isValid(string $s): bool {
        $brackets = [];
        for ($i = 0; $i < strlen($s); $i++) {
            $item = $s[$i];
            if (isset(self::CLOSING_BRACKETS[$item])) {
                $last = array_pop($brackets);
                if ($last !== self::CLOSING_BRACKETS[$item]) {
                    return false;
                }

                continue;
            }

            $brackets[] = $item;
        }

        return count($brackets) === 0;
    }
}