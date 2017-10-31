package regressiontests

import "testing"

func TestNakedret(t *testing.T) {
	t.Parallel()
	source := `package test

	func hi() (uint32, error) {
//
//
//
//
//
//
		return
	}
`
	expected := Issues{
		{Linter: "nakedret", Severity: "warning", Path: "test.go", Line: 3, Col: 16, Message: "struct of size 24 could be 16"},
	}
	ExpectIssues(t, "nakedret", source, expected)
}
