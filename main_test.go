package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMathOperations1(t *testing.T) {
	numbersA := []int{1, 2, 3}
	numbersB := []int{1, 2, 3}
	result := []int{2, 0, 9}
	operations := []string{"+", "-", "*"}

	sliceReturned := MathOperations(numbersA, numbersB, operations)

	assert.Equal(t, result, sliceReturned)
}

func TestMathOperations2(t *testing.T) {
	numbersA := []int{1, 2, 3, 3}
	numbersB := []int{1, 2, 3, 3}
	result := []int{2, 0, 9, 6}
	operations := []string{"+", "-", "*", "+"}

	sliceReturned := MathOperations(numbersA, numbersB, operations)

	assert.Equal(t, result, sliceReturned)
}
