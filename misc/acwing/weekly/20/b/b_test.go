// Code generated by copypasta/template/acwing/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [b]")
	testCases := [][2]string{
		{
			`2 0 0
1 2
2 3`,
			`2`,
		},
		{
			`2 1 0
1 2
2 2`,
			`0`,
		},
		{
			`2 5 7
3 4
14 4`,
			`1`,
		},
		
	}
	target := 0 // -1
	testutil.AssertEqualStringCase(t, testCases, target, run)
}
