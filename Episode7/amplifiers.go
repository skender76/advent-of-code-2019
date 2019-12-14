package main

import (
	"fmt"
)

type amplifier struct {
	name			   string
	program            []int
	phase_setting      int
	halted             bool
	instructionPointer int
	input_ptr		   int
}

func (r *amplifier)calcOutput(input int) int{
	var output int = 0
	var input_val = []int{r.phase_setting,input}

	if !r.halted {
		r.halted,output = r.runIntcodeComputer(input_val)
	} else {
		fmt.Println("Input for halted amplifier")
	}

	return output
}

func (r *amplifier)runIntcodeComputer(input_val []int) (bool, int) {
	var halt_reached = false
	var result = 0

	for  r.instructionPointer < len(r.program) {
		var targetPos = 0
		var value = 0

		var instruction, mode_first, mode_second, out_mode = readOpcode(r.program[r.instructionPointer])

		if instruction == ADD {
			targetPos, value = calcSum(r.program, r.instructionPointer, mode_first, mode_second, out_mode)
			r.program[targetPos] = value
			r.instructionPointer += OFFSET_SUM
		} else if instruction == MUL {
			targetPos, value = calcMul(r.program, r.instructionPointer, mode_first, mode_second, out_mode)
			r.program[targetPos] = value
			r.instructionPointer += OFFSET_MUL
		} else if instruction == HALT {
			halt_reached = true
			r.instructionPointer += OFFSET_HALT
			break
		} else if instruction == SAVE {
			var targetPos = calcPos(r.program, r.instructionPointer+1, mode_first)
			r.program[targetPos] = input_val[r.input_ptr]
			if r.input_ptr < (len(input_val) - 1) {
				r.input_ptr++
			}

			r.instructionPointer += OFFSET_SAVE
		} else if instruction == OUT {
			var targetPos = calcPos(r.program, r.instructionPointer+1, mode_first)
			result=  r.program[targetPos]
			r.instructionPointer += OFFSET_OUT
			break
		} else if instruction == EQUAL {
			targetPos, value = calcEqual(r.program, r.instructionPointer, mode_first, mode_second, out_mode)
			r.program[targetPos] = value
			r.instructionPointer += OFFSET_EQUAL
		} else if instruction == JUMP_IF_FALSE {
			r.instructionPointer = calcJumpIfFalsePos(r.program, r.instructionPointer, mode_first, mode_second)
		} else if instruction == JUMP_IF_TRUE {
			r.instructionPointer = calcJumpIfTruePos(r.program, r.instructionPointer, mode_first, mode_second)
		} else if instruction == LESS_THAN {
			targetPos, value = calcLessThan(r.program, r.instructionPointer, mode_first, mode_second, out_mode)
			r.program[targetPos] = value
			r.instructionPointer += OFFSET_LESS_THAN
		}else {
			fmt.Println("Wrong...")
			break
		}
	}

	return halt_reached, result
}


func calcAmpOutputWithFeedback(settingSequence, input []int) int {

	var A = amplifier{"A",clone(input), settingSequence[0], false, 0,0}
	var B = amplifier{"B",clone(input), settingSequence[1], false,0,0 }
	var C = amplifier{"C",clone(input), settingSequence[2], false,0,0}
	var D = amplifier{"D",clone(input), settingSequence[3], false,0,0}
	var E = amplifier{"E",clone(input), settingSequence[4], false,0,0}

	var halted = false
	var result = 0

	for !halted {

		var curr_result = result
		curr_result = A.calcOutput(curr_result)
		curr_result = B.calcOutput(curr_result)
		curr_result = C.calcOutput(curr_result)
		curr_result = D.calcOutput(curr_result)
		curr_result = E.calcOutput(curr_result)
		if E.halted {
			break
		} else {
			result = curr_result
		}
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

func calcWaterfallAmpOutput(settingSequence, input []int) int {

	var A = amplifier{"A",clone(input), settingSequence[0], false,0,0}
	var B = amplifier{"B",clone(input), settingSequence[1], false,0,0}
	var C = amplifier{"C",clone(input), settingSequence[2], false,0,0}
	var D = amplifier{"D",clone(input), settingSequence[3], false,0,0}
	var E = amplifier{"E",clone(input), settingSequence[4], false,0,0}

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
