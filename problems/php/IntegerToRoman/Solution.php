<?php
namespace Problem\IntegerToRoman;

class Solution {

    /**
     * @param Integer $num
     * @return String
     */
    function intToRoman($num) {
        $valueRomans = [
            1000 => 'M',
            900 => 'CM',
            500 => 'D',
            400 => 'CD',
            100 => 'C',
            90 => 'XC',
            50 => 'L',
            40 => 'XL',
            10 => 'X',
            9 => 'IX',
            5 => 'V',
            4 => 'IV',
            1 => 'I',
        ];

        $res = "";
        foreach ($valueRomans as $value => $roman) {
            while ($num >= $value) {
                $num -= $value;
                $res .= $roman;
            }
        }

        return $res;
    }
}