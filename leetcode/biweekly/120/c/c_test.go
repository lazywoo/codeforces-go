// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_c(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, incremovableSubarrayCount, "c.txt", 0); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/biweekly-contest-120/problems/count-the-number-of-incremovable-subarrays-ii/
// https://leetcode.cn/problems/count-the-number-of-incremovable-subarrays-ii/