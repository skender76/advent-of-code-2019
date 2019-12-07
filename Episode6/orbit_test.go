package main

import "testing"

func TestOrbit(t *testing.T) {

	var HEAD = Planet{"COM", nil}
	var input = []string{"COM)B","B)C","C)D","D)E","E)F","B)G","G)H","D)I","E)J","J)K","K)L"}

	calculateOrbit(&HEAD, input)

	var result = -1
	result = count_orbit(HEAD, result)

	if result != 42 {
		t.Errorf("Value %d but expected %d", result, 42) // to indicate test failed
	}

}