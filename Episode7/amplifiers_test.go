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

func TestCalcOutputValueWithFeedback(t *testing.T) {
	var settingSequence = []int{9,8,7,6,5}
	var input = []int {3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5}

	var result = calcAmpOutputWithFeedback(settingSequence, input)

	if result != 139629729 {
		t.Errorf("Value %d but expected %d", result , 139629729) // to indicate test failed
	}
}

func TestCalcOutputValueWithFeedback2(t *testing.T) {
	var settingSequence = []int{9,7,8,5,6}
	var input = []int {3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,
		-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,
		53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10}

	var result = calcAmpOutputWithFeedback(settingSequence, input)

	if result != 18216 {
		t.Errorf("Value %d but expected %d", result , 18216) // to indicate test failed
	}
}


