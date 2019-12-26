package main

import "testing"

func TestRemove(t *testing.T) {

	var input = []int {0,1,2,3,4}

	var result = remove(input, 2)

	if len(input) != 5 {
		t.Errorf("Value %d but expected %d", len(input) , 5)
	}

	if len(result) != 4 {
		t.Errorf("Value %d but expected %d", len(result) , 4)
	}

	if !testEq(result,[]int {0,1,3,4}) {
		t.Errorf("Value %d but expected %d", result , []int {0,1,3,4}) // to indicate test failed
	}
}

func TestRemove2(t *testing.T) {

	var input = []int {0,1,2,3,4}
	var result = make([]int, len(input))
	copy(result, input)
	result = remove(result, 2)

	if !testEq(input,[]int {0,1,2,3,4}) {
		t.Errorf("Input %d but expected %d", result , []int {0,1,2,3,4}) // to indicate test failed
	}

	if !testEq(result,[]int {0,1,3,4}) {
		t.Errorf("Result %d but expected %d", result , []int {0,1,3,4}) // to indicate test failed
	}
}

func testEq(a, b []int) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func testEqFloat(t *testing.T) {

	x:= 5.0
	y:=6.0

	if equal(x,y) {
		t.Errorf("Error %f is different from %f", x, y)
	}

	y = 5.0

	if !equal(x,y) {
		t.Errorf("Error %f is equal to %f", x, y)
	}


}