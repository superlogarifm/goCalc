package calculate_test

import (
	"CalcServer/calculate"
	"testing"
)

func TestCalc(t *testing.T) {
	testCasesSuccess := [...]struct {
		casename       string
		expression     string
		expectedResult float64
	}{
		{"easy", "2+2", 4},
		{"priority", "(2+2)*2", 8},
		{"priority2", "2+2*2", 6},
		{"div", "4/2", 2},
		{"zero", "0/5", 0},
	}

	for _, testCase := range testCasesSuccess[:] {
		t.Run(testCase.casename, func(t *testing.T) {
			val, err := calculate.Calc(testCase.expression)
			if err != nil {
				t.Fatalf("Error %s", testCase.expression)
			}
			if val != testCase.expectedResult {
				t.Fatalf("expected %f, got %f", testCase.expectedResult, val)
			}
		})
	}

	testCasesFail := [...]struct {
		casename      string
		expression    string
		expectedError error
	}{
		{"invalid operator", "41-41**41", calculate.ErrInvalidExpression},
		{"parentheses", "((5+2)-*(4", calculate.ErrInvalidExpression},
		{"empty", "", calculate.ErrInvalidExpression},
		{"zero division", "5/0", calculate.ErrZeroDivision},
		{"Inc input", "a+5", calculate.ErrIncorrectInput},
	}

	for _, testCase := range testCasesFail[:] {
		t.Run(testCase.casename, func(t *testing.T) {
			_, err := calculate.Calc(testCase.expression)
			if err == nil {
				t.Fatalf("expected error: %s, but got nil", testCase.expectedError)
			}
		})
	}
}
