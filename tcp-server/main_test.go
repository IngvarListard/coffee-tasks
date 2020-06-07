package main

import (
	"fmt"
	"testing"
)

var expression, expectedResult string

func testResolveExpression(t *testing.T, exp, exRes string) {
	if res := resolveExpression(exp); res != exRes {
		fmt.Println(res == exRes)
		t.Errorf("expected: %v, got %v", exRes, res)
	}
}

func TestResolveExpressionSimpleCases(t *testing.T) {

	expression = "5000 + 2222"
	expectedResult = "Result: 7222"
	testResolveExpression(t, expression, expectedResult)

	expression = "5000 / 3"
	expectedResult = "Result: 1666.6666666666667"
	testResolveExpression(t, expression, expectedResult)

	expression = "2732 * 221"
	expectedResult = "Result: 603772"
	testResolveExpression(t, expression, expectedResult)

	expression = "87124 - 77124"
	expectedResult = "Result: 10000"
	testResolveExpression(t, expression, expectedResult)
}

func TestResolveExpressionFailedCases(t *testing.T) {
	expression = "TEST"
	expectedResult = "error: wrong expression format"
	testResolveExpression(t, expression, expectedResult)

	expression = "55+27"
	expectedResult = "error: wrong expression format"
	testResolveExpression(t, expression, expectedResult)

	expression = ""
	expectedResult = "error: wrong expression format"
	testResolveExpression(t, expression, expectedResult)
}
