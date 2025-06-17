<?php
namespace Problem\MergeTwoSortedLists;

use PHPUnit\Framework\TestCase;

class SolutionTest extends TestCase {
    public function testMergeTwoLists(): void {
        $solution = new Solution();

        $this->assertEquals(
            new ListNode(1,
                new ListNode(1,
                    new ListNode(2,
                        new ListNode(3,
                            new ListNode(4,
                                new ListNode(4)
            ))))),
            $solution->mergeTwoLists(
                new ListNode(1,
                    new ListNode(2,
                        new ListNode(4)
                )),
                new ListNode(1,
                    new ListNode(3,
                        new ListNode(4)
                ))
            )
        );

        $this->assertEquals(null, $solution->mergeTwoLists(null,null));

        $this->assertEquals(new ListNode(0), $solution->mergeTwoLists(null, new ListNode(0)));
    }
}