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

//type rect struct {
//	width int
//	height int
//}
//
//func (r *rect) area() int {
//	return r.width * r.height
//}

func amplifier_output(program []int, phase_setting, input int) int {
	var input_val = []int{phase_setting,input}
	var _,result = runIntcodeComputer(input_val,  program )
	return result
}

func calcWaterfallAmpOutput(settingSequence, input []int) int {

	var result = amplifier_output(input,settingSequence[0],0)
	result = amplifier_output(input,settingSequence[1],result)
	result = amplifier_output(input,settingSequence[2],result)
	result = amplifier_output(input,settingSequence[3],result)
	result = amplifier_output(input,settingSequence[4],result)

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

func calcMaxThrustherSignalWithFeedback(sequence, input []int) int {
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

func main() {
	var sequence = []int{5,6,7,8,9}
	var input = []int{3,8,1001,8,10,8,105,1,0,0,21,34,59,68,89,102,183,264,345,426,99999,3,9,102,5,9,9,1001,9,5,9,4,9,99,3,9,101,3,9,9,1002,9,5,9,101,5,9,9,1002,9,3,9,1001,9,5,9,4,9,99,3,9,101,5,9,9,4,9,99,3,9,102,4,9,9,101,3,9,9,102,5,9,9,101,4,9,9,4,9,99,3,9,1002,9,5,9,1001,9,2,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,99,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,99}
	var max = calcMaxThrustherSignalWithFeedback(sequence, input)

	fmt.Println(max)
}
