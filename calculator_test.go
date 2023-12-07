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

func TestDivide(t *testing.T) {
	result, err := Divide(6, 3)
	assert.NoError(t, err, "an unexpected error occured: %v", err)
	assert.Equal(t, 2, result, "they should be equal")

	// if err != nil {
	// 	t.Errorf("Divide(6, 3) failed. Got error %s", err.Error())
	// }
	// if result != 2 {
	// 	t.Errorf("Divide(6, 3) failed. Expected 2, got %d", result)
	// }
}
