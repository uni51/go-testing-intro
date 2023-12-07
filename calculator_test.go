package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	expected := 3
	if result != expected {
		t.Errorf("Add(1, 2) failed. Expected 3, got %d", result)
	}
}

type devideTestCase struct {
	a, b           int
	expectedResult int
	expectError    bool
}

func TestDivide(t *testing.T) {
	testCases := []devideTestCase{
		{a: 6, b: 3, expectedResult: 2, expectError: false},
		{a: 6, b: 0, expectedResult: 0, expectError: true},
	}

	for _, tc := range testCases {
		result, err := Divide(tc.a, tc.b)
		if tc.expectError {
			assert.Error(t, err, "an unexpected error occured: %v", err)
		} else {
			assert.NoError(t, err, "an unexpected error occured: %v", err)
			assert.Equal(t, tc.expectedResult, result, "they should be equal")
		}

		// if tc.expectError {
		// 	if err == nil {
		// 		t.Errorf("Divide(%d, %d) failed. Expected an error, got nil", tc.a, tc.b)
		// 	}
		// } else {
		// 	if err != nil {
		// 		t.Errorf("Divide(%d, %d) failed. Expected no error, got %s", tc.a, tc.b, err.Error())
		// 	}
		// }
		// if result != tc.expectedResult {
		// 	t.Errorf("Divide(%d, %d), expected %d, got %d", tc.a, tc.b, tc.expectedResult, result)
		// }
	}

	// result, err := Divide(6, 3)
	// assert.NoError(t, err, "an unexpected error occured: %v", err)
	// assert.Equal(t, 2, result, "they should be equal")

	// if err != nil {
	// 	t.Errorf("Divide(6, 3) failed. Got error %s", err.Error())
	// }
	// if result != 2 {
	// 	t.Errorf("Divide(6, 3) failed. Expected 2, got %d", result)
	// }
}
