// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [e]")
	examples := [][]string{
		{
			`["W","N","ES","W"]`,
			`2`,
		},
		{
			`["NS","WE","SE","EW"]`,
			`3`,
		},

	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, trafficCommand, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/season/2021-fall/problems/Y1VbOX/
