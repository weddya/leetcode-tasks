<?php
namespace Problem\LongestCommonPrefix;

class Solution {

    /**
     * @param String[] $strs
     * @return String
     */
    function longestCommonPrefix($strs) {
        $prefix = "";
        foreach ($strs as $strInd => $str) {
            if ($strInd == 0) {
                $prefix = $str;
            }

            for ($i = 0; $i < max(strlen($prefix), strlen($str)); $i++) {
                if (!isset($prefix[$i])
                    || !isset($str[$i])
                    || $prefix[$i] !== $str[$i]
                ) {
                    $prefix = substr($prefix, 0, $i);
                    break;
                }
            }
        }

        return $prefix;
    }
}