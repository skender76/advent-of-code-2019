package main

import "fmt"

type amplifier struct {
	program []int
	phase_setting int
	halted bool
}

func (r *amplifier)calcOutput(input int) int{
	var input_val = []int{r.phase_setting,input}
	var halted,output = runIntcodeComputer(input_val,  r.program )
	r.halted = halted
	return output
}

func amplifier_output(program []int, phase_setting, input int) int {
	var input_val = []int{phase_setting,input}
	var _,result = runIntcodeComputer(input_val,  program )
	return result
}

func calcWaterfallAmpOutput(settingSequence, input []int) int {

	var A = amplifier{input, settingSequence[0], false}
	var B = amplifier{input, settingSequence[1], false}
	var C = amplifier{input, settingSequence[2], false}
	var D = amplifier{input, settingSequence[3], false}
	var E = amplifier{input, settingSequence[4], false}

	var halted = false
	var result = A.calcOutput(0)

	for !halted {
		result = B.calcOutput(result)
		result = C.calcOutput(result)
		result = D.calcOutput(result)
		result = E.calcOutput(result)

		halted = A.halted && B.halted && C.halted && D.halted && E.halted
	}

	return result
}

func calcMaxThrustherSignal(sequence, input []int) int {
	var phaseSettings = calcPermutationWithoutRepeatingNumber(sequence)
	var maxSignal = 0

	for index:=0;index<len(phaseSettings);index++ {

		var thrustherSignal = calcWaterfallAmpOutput(phaseSettings[index], input)

		if thrustherSignal > maxSignal {
			maxSignal = thrustherSignal
		}
	}

	return maxSignal
}

func calcAmpOutputWithFeedback(settingSequence, input []int) int {

	var A = amplifier{input, settingSequence[0], false}
	var B = amplifier{input, settingSequence[1], false}
	var C = amplifier{input, settingSequence[2], false}
	var D = amplifier{input, settingSequence[3], false}
	var E = amplifier{input, settingSequence[4], false}

	var result = 0


	result = A.calcOutput(result)
	result = B.calcOutput(result)
	result = C.calcOutput(result)
	result = D.calcOutput(result)
	result = E.calcOutput(result)

	return result
}

func calcMaxThrustherSignalWithFeedback(sequence, input []int) int {
	var phaseSettings = calcPermutationWithoutRepeatingNumber(sequence)
	var maxSignal = 0

	for index:=0;index<len(phaseSettings);index++ {

		var thrustherSignal = calcAmpOutputWithFeedback(phaseSettings[index], input)

		if thrustherSignal > maxSignal {
			maxSignal = thrustherSignal
		}
	}

	return maxSignal
}

func main() {
	var sequence = []int{5,6,7,8,9}
	var input = []int{3,8,1001,8,10,8,105,1,0,0,21,34,59,68,89,102,183,264,345,426,99999,3,9,102,5,9,9,1001,9,5,9,4,9,99,3,9,101,3,9,9,1002,9,5,9,101,5,9,9,1002,9,3,9,1001,9,5,9,4,9,99,3,9,101,5,9,9,4,9,99,3,9,102,4,9,9,101,3,9,9,102,5,9,9,101,4,9,9,4,9,99,3,9,1002,9,5,9,1001,9,2,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,99,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,99}
	var max = calcMaxThrustherSignalWithFeedback(sequence, input)

	fmt.Println(max)
}
