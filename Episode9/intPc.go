package main

import "fmt"

const ADD int = 1
const MUL int = 2
const SAVE int = 3
const OUT int = 4
const HALT int = 99
const JUMP_IF_TRUE int = 5
const JUMP_IF_FALSE int = 6
const LESS_THAN int = 7
const EQUAL int = 8
const ADJUST_RELATIVE_BASE int = 9

const OFFSET_MUL int = 4
const OFFSET_SUM int = 4
const OFFSET_HALT int = 1
const OFFSET_SAVE int = 2
const OFFSET_OUT int = 2

const OFFSET_JUMP_IF_TRUE int = 3
const OFFSET_JUMP_IF_FALSE int = 3
const OFFSET_LESS_THAN int = 4
const OFFSET_EQUAL int = 4
const OFFSET_ADJUST_RELATIVE_BASE int = 2

const POSITION_MODE int = 0
const IMMEDIATE_MODE int = 1
const RELATIVE_MODE int = 2


type IntCodeComputer struct {
	program            []int
	halted             bool
	instructionPointer int
	input_ptr		   int
	relative_base		int
}

func NewIntCodeComputer(program []int) *IntCodeComputer {
	p := new(IntCodeComputer)
	p.program = cloneExt(program, 5000)
	p.halted = false
	p.instructionPointer = 0
	p.input_ptr = 0
	p.relative_base = 0

	return p
}

func (r *IntCodeComputer)execProgram(input_val []int) int {
	var result = 0

	for  r.instructionPointer < len(r.program) {
		var targetPos = 0
		var value = 0

		var instruction, mode_first, mode_second, out_mode = r.readOpcode(r.program[r.instructionPointer])

		if instruction == ADD {
			targetPos, value = r.calcSum(r.program, r.instructionPointer, mode_first, mode_second, out_mode)
			r.program[targetPos] = value
			r.instructionPointer += OFFSET_SUM
		} else if instruction == MUL {
			targetPos, value = r.calcMul(r.program, r.instructionPointer, mode_first, mode_second, out_mode)
			r.program[targetPos] = value
			r.instructionPointer += OFFSET_MUL
		} else if instruction == HALT {
			r.halted = true
			r.instructionPointer += OFFSET_HALT
			break
		} else if instruction == SAVE {
			var targetPos = r.calcPos(r.program, r.instructionPointer+1, mode_first)
			r.program[targetPos] = input_val[r.input_ptr]
			if r.input_ptr < (len(input_val) - 1) {
				r.input_ptr++
			}
			r.instructionPointer += OFFSET_SAVE
		} else if instruction == OUT {
			var targetPos = r.calcPos(r.program, r.instructionPointer+1, mode_first)
			result=  r.program[targetPos]
			r.instructionPointer += OFFSET_OUT
			break
		} else if instruction == EQUAL {
			targetPos, value = r.calcEqual(r.program, r.instructionPointer, mode_first, mode_second, out_mode)
			r.program[targetPos] = value
			r.instructionPointer += OFFSET_EQUAL
		} else if instruction == JUMP_IF_FALSE {
			r.instructionPointer = r.calcJumpIfFalsePos(r.program, r.instructionPointer, mode_first, mode_second)
		} else if instruction == JUMP_IF_TRUE {
			r.instructionPointer = r.calcJumpIfTruePos(r.program, r.instructionPointer, mode_first, mode_second)
		} else if instruction == LESS_THAN {
			targetPos, value = r.calcLessThan(r.program, r.instructionPointer, mode_first, mode_second, out_mode)
			r.program[targetPos] = value
			r.instructionPointer += OFFSET_LESS_THAN
		} else if instruction == ADJUST_RELATIVE_BASE {
			r.relative_base = r.calcRelativeBasePos(r.program, r.instructionPointer+1, mode_first)
			r.instructionPointer += OFFSET_ADJUST_RELATIVE_BASE
		}else {
			fmt.Println("Wrong...")
			break
		}
	}

	return result
}



func (r *IntCodeComputer)convertInsCodeToStr( instCode int ) string {
	var codeStr = ""
	var instCodeExt,_,_,_ = r.readOpcode(instCode)
	if instCodeExt == ADD {
		codeStr = "Add"
	} else if instCodeExt == MUL {
		codeStr = "Mul"
	} else if instCodeExt == SAVE {
		codeStr = "Save"
	} else if instCodeExt == OUT {
		codeStr = "Out"
	} else if instCodeExt == HALT {
		codeStr = "Halt"
	} else if instCodeExt == JUMP_IF_TRUE {
		codeStr = "Jump If True"
	} else if instCodeExt == JUMP_IF_FALSE {
		codeStr = "Jump If False"
	} else if instCodeExt == LESS_THAN {
		codeStr = "Less Than"
	} else if instCodeExt == EQUAL {
		codeStr = "Equal"
	}

	return codeStr
}

func (r *IntCodeComputer)readOpcode( input int ) (int, int, int, int) {

	var op_code = input % 100
	var mode_first = (input/100) % 10
	var mode_second = (input/1000) % 10
	var mode_third = (input/10000) % 10

	return op_code, mode_first, mode_second, mode_third
}

func (r *IntCodeComputer)calcRelativeBasePos(input []int, pos, mode int) int {
	var result_pos = 0

	if mode == POSITION_MODE {
		result_pos = input[pos]
	} else if mode == IMMEDIATE_MODE {
		result_pos = pos
	} else if mode == RELATIVE_MODE {
		result_pos = r.relative_base + input[pos]
	}

	result_pos = r.relative_base + r.program[result_pos]

	return result_pos
}

func (r *IntCodeComputer)calcPos(input []int, pos, mode int) int {
	var result_pos = 0

	if mode == POSITION_MODE {
		result_pos = input[pos]
	} else if mode == IMMEDIATE_MODE {
		result_pos = pos
	} else if mode == RELATIVE_MODE {
		result_pos = r.relative_base + input[pos]
	}

	return result_pos
}

func (r *IntCodeComputer)calcJumpIfFalsePos(input []int, op_code_pos, mode_first, mode_second int) int  {
	var firstArgPos = r.calcPos(input, op_code_pos+1, mode_first)
	var secondArgPos = r.calcPos(input, op_code_pos+2, mode_second)
	var new_pos = op_code_pos
	var offset = OFFSET_JUMP_IF_FALSE
	if input[firstArgPos] == 0 {
		new_pos = input[secondArgPos]
		offset = 0
	} else {
		new_pos += offset
	}
	return new_pos
}

func (r *IntCodeComputer)calcJumpIfTruePos(input []int, op_code_pos, mode_first, mode_second int) int {
	var firstArgPos = r.calcPos(input, op_code_pos+1, mode_first)
	var secondArgPos = r.calcPos(input, op_code_pos+2, mode_second)
	var new_pos = op_code_pos
	var offset = OFFSET_JUMP_IF_TRUE
	if input[firstArgPos] != 0 {
		new_pos = input[secondArgPos]
	} else {
		new_pos += offset
	}
	return new_pos
}

func (r *IntCodeComputer)calcEqual(input []int, op_code_pos, mode_first, mode_second, out_mode int) (int, int) {
	var firstArgPos = r.calcPos(input, op_code_pos+1, mode_first)
	var secondArgPos = r.calcPos(input, op_code_pos+2, mode_second)
	var targetPos = r.calcPos(input, op_code_pos+3, out_mode)
	var value = 0
	if input[firstArgPos] == input[secondArgPos] {
		value = 1
	}
	return targetPos, value
}

func (r *IntCodeComputer)calcLessThan(input []int, op_code_pos, mode_first, mode_second, out_mode int) (int, int) {
	var firstArgPos = r.calcPos(input, op_code_pos+1, mode_first)
	var secondArgPos = r.calcPos(input, op_code_pos+2, mode_second)
	var targetPos = r.calcPos(input, op_code_pos+3, out_mode)
	var value = 0
	if input[firstArgPos] < input[secondArgPos] {
		value = 1
	}
	return targetPos, value
}

func (r *IntCodeComputer)calcMul(input []int, op_code_pos, mode_first, mode_second, out_mode int) (int, int) {
	var firstArgPos = r.calcPos(input, op_code_pos+1, mode_first)
	var secondArgPos = r.calcPos(input, op_code_pos+2, mode_second)
	var targetPos = r.calcPos(input, op_code_pos+3, out_mode)
	var value = input[firstArgPos] * input[secondArgPos]

	return targetPos, value
}

func (r *IntCodeComputer)calcSum(input []int, op_code_pos, mode_first, mode_second, out_mode int) (int, int) {
	var firstArgPos = r.calcPos(input, op_code_pos+1, mode_first)
	var secondArgPos = r.calcPos(input, op_code_pos+2, mode_second)
	var targetPos = r.calcPos(input, op_code_pos+3, out_mode)

	var value = input[firstArgPos] + input[secondArgPos]

	return targetPos, value
}

