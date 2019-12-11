package main

import "testing"

func TestCalcOutputValue(t *testing.T) {
	var settingSequence = []int{4,3,2,1,0}
	var input = []int {3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0}

	var result = calcWaterfallAmpOutput(settingSequence, input)

	if result != 43210 {
		t.Errorf("Value %d but expected %d", result , 43210) // to indicate test failed
	}
}

func TestCalcOutputValue2(t *testing.T) {
	var settingSequence = []int{0,1,2,3,4}
	var input = []int {3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0}

	var result = calcWaterfallAmpOutput(settingSequence, input)

	if result != 54321 {
		t.Errorf("Value %d but expected %d", result , 54321) // to indicate test failed
	}
}

func TestCalcOutputValue3(t *testing.T) {
	var settingSequence = []int{1,0,4,3,2}
	var input = []int {3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0}

	var result = calcWaterfallAmpOutput(settingSequence, input)

	if result != 65210 {
		t.Errorf("Value %d but expected %d", result , 65210) // to indicate test failed
	}
}

func TestCalcMaxThrustherSignal(t *testing.T) {
	var sequence = []int{0,1,2,3,4,5}
	var input = []int {3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0}

	var result = calcMaxThrustherSignal(sequence,input)

	if result != 65210 {
		t.Errorf("Value %d but expected %d", result , 65210) // to indicate test failed
	}
}

func TestCalcMaxThrustherSignal2(t *testing.T) {
	var sequence = []int{0,1,2,3,4,5}
	var input = []int {3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0}

	var result = calcMaxThrustherSignal(sequence, input)

	if result != 54321 {
		t.Errorf("Value %d but expected %d", result , 54321) // to indicate test failed
	}
}



