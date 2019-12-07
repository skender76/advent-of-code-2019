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

func TestTrivialJump(t *testing.T) {

	var HEAD = Planet{"COM", nil}
	var input = []string{"COM)YOU","COM)SAN"}

	calculateOrbit(&HEAD, input)

	var result = calc_orbital_transfers(HEAD, "YOU", "SAN")

	if result != 0 {
		t.Errorf("Value %d but expected %d", result, 0) // to indicate test failed
	}

}

func TestSimpleJump(t *testing.T) {

	var HEAD = Planet{"COM", nil}
	var input = []string{"COM)B","COM)D","B)YOU","COM)SAN","COM)C"}

	calculateOrbit(&HEAD, input)

	var result = calc_orbital_transfers(HEAD, "YOU", "SAN")

	if result != 1 {
		t.Errorf("Value %d but expected %d", result, 1) // to indicate test failed
	}

}

func TestComplexJump(t *testing.T) {

	var HEAD = Planet{"COM", nil}
	var input = []string{"COM)B","COM)D","B)YOU","D)SAN","COM)C"}

	calculateOrbit(&HEAD, input)

	var result = calc_orbital_transfers(HEAD, "YOU", "SAN")

	if result != 2 {
		t.Errorf("Value %d but expected %d", result, 2) // to indicate test failed
	}

}

func TestSampleJump(t *testing.T) {

	var HEAD = Planet{"COM", nil}
	var input = []string{"COM)B","B)C","C)D","D)E","E)F","B)G","G)H","D)I","E)J","J)K","K)L","K)YOU","I)SAN"}

	calculateOrbit(&HEAD, input)

	var result = calc_orbital_transfers(HEAD, "YOU", "SAN")

	if result != 4 {
		t.Errorf("Value %d but expected %d", result, 4) // to indicate test failed
	}

}
