package main

import "testing"

func TestPermutation(t *testing.T) {
	var result = calcPermutationWithRepeatingNumber()

	if len(result) != 3125 {
		t.Errorf("Value %d but expected %d", len(result) , 3125) // to indicate test failed
	}
}

func TestPermutationWithoutRepeatingNumber(t *testing.T) {
	var sequence = []int{0,1,2,3,4,5}
	var result = calcPermutationWithoutRepeatingNumber(sequence)

	if len(result) != 120 {
		t.Errorf("Value %d but expected %d", len(result) , 120) // to indicate test failed
	}
}
