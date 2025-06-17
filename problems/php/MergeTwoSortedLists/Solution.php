<?php
namespace Problem\MergeTwoSortedLists;

// Definition for a singly-linked list.
class ListNode {
    public int $val = 0;
    public ListNode|null $next = null;

    function __construct($val = 0, $next = null) {
        $this->val = $val;
        $this->next = $next;
    }
}

class Solution {

    /**
     * @param ?ListNode $list1
     * @param ?ListNode $list2
     * @return ?ListNode
     */
    function mergeTwoLists(ListNode|null $list1, ListNode|null $list2): ListNode|null {
        if ($list1 == null) return $list2;
        if ($list2 == null) return $list1;

        if ($list1->val < $list2->val) {
            $list1->next = $this->mergeTwoLists($list1->next, $list2);
            return $list1;
        } else {
            $list2->next = $this->mergeTwoLists($list2->next, $list1);
            return $list2;
        }
    }
}