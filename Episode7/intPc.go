package main

import (
	"fmt"
)

const ADD int = 1
const MUL int = 2
const SAVE int = 3
const OUT int = 4
const HALT int = 99
const JUMP_IF_TRUE int = 5
const JUMP_IF_FALSE int = 6
const LESS_THAN int = 7
const EQUAL int = 8

const OFFSET_MUL int = 4
const OFFSET_SUM int = 4
const OFFSET_HALT int = 1
const OFFSET_SAVE int = 2
const OFFSET_OUT int = 2

const OFFSET_JUMP_IF_TRUE int = 3
const OFFSET_JUMP_IF_FALSE int = 3
const OFFSET_LESS_THAN int = 4
const OFFSET_EQUAL int = 4

const POSITION_MODE int = 0
const IMMEDIATE_MODE int = 1

func readOpcode( input int ) (int, int, int, int) {

	var op_code = input % 100
	var mode_first = (input/100) % 10
	var mode_second = (input/1000) % 10
	var mode_third = (input/10000) % 10

	return op_code, mode_first, mode_second, mode_third
}

func runIntcodeComputer(input_val []int, input []int) (bool, int) {
	var halt_reached = false
	var result = -1
	var nextOpCodeOffset = 4
	var input_ptr = 0

	for op_code_pos := 0; op_code_pos < len(input); op_code_pos += nextOpCodeOffset {
		var targetPos = 0
		var value = 0

		var opcode, mode_first, mode_second, out_mode = readOpcode(input[op_code_pos])

		if opcode == ADD {
			targetPos, value = calcSum(input, op_code_pos, mode_first, mode_second, out_mode)
			input[targetPos] = value
			nextOpCodeOffset = OFFSET_SUM
		} else if opcode == MUL {
			targetPos, value = calcMul(input, op_code_pos, mode_first, mode_second, out_mode)
			input[targetPos] = value
			nextOpCodeOffset = OFFSET_MUL
		} else if opcode == HALT {
			halt_reached = true
			nextOpCodeOffset = OFFSET_HALT
			break
		} else if opcode == SAVE {
			var targetPos = calcPos(input, op_code_pos+1, mode_first)
			input[targetPos] = input_val[input_ptr]

			if input_ptr < (len(input_val) - 1) {
				input_ptr++
			}

			nextOpCodeOffset = OFFSET_SAVE
		} else if opcode == OUT {
			var targetPos = calcPos(input, op_code_pos+1, mode_first)
			result=  input[targetPos]
			nextOpCodeOffset = OFFSET_OUT
			break
		} else if opcode == EQUAL {
			targetPos, value = calcEqual(input, op_code_pos, mode_first, mode_second, out_mode)
			input[targetPos] = value
			nextOpCodeOffset = OFFSET_EQUAL
		} else if opcode == JUMP_IF_FALSE {
			op_code_pos, nextOpCodeOffset = calcJumpIfFalsePos(input, op_code_pos, mode_first, mode_second)
		} else if opcode == JUMP_IF_TRUE {
			op_code_pos, nextOpCodeOffset = calcJumpIfTruePos(input, op_code_pos, mode_first, mode_second)
		} else if opcode == LESS_THAN {
			targetPos, value = calcLessThan(input, op_code_pos, mode_first, mode_second, out_mode)
			input[targetPos] = value
			nextOpCodeOffset = OFFSET_LESS_THAN
		}else {
			fmt.Println("Wrong...")
			break
		}
	}

	return halt_reached, result
}

func calcPos(input []int, pos, mode int) int {
	var result_pos = 0

	if mode == POSITION_MODE {
		result_pos = input[pos]
	} else if mode == IMMEDIATE_MODE {
		result_pos = pos
	}

	return result_pos
}

func calcJumpIfFalsePos(input []int, op_code_pos, mode_first, mode_second int) (int,int)  {
	var firstArgPos = calcPos(input, op_code_pos+1, mode_first)
	var secondArgPos = calcPos(input, op_code_pos+2, mode_second)
	var new_pos = op_code_pos
	var offset = OFFSET_JUMP_IF_FALSE
	if input[firstArgPos] == 0 {
		new_pos = input[secondArgPos]
		offset = 0
	}
	return new_pos, offset
}

func calcJumpIfTruePos(input []int, op_code_pos, mode_first, mode_second int) (int,int) {
	var firstArgPos = calcPos(input, op_code_pos+1, mode_first)
	var secondArgPos = calcPos(input, op_code_pos+2, mode_second)
	var new_pos = op_code_pos
	var offset = OFFSET_JUMP_IF_TRUE
	if input[firstArgPos] != 0 {
		new_pos = input[secondArgPos]
		offset = 0
	}
	return new_pos, offset
}

func calcEqual(input []int, op_code_pos, mode_first, mode_second, out_mode int) (int, int) {
	var firstArgPos = calcPos(input, op_code_pos+1, mode_first)
	var secondArgPos = calcPos(input, op_code_pos+2, mode_second)
	var targetPos = calcPos(input, op_code_pos+3, out_mode)
	var value = 0
	if input[firstArgPos] == input[secondArgPos] {
		value = 1
	}
	return targetPos, value
}

func calcLessThan(input []int, op_code_pos, mode_first, mode_second, out_mode int) (int, int) {
	var firstArgPos = calcPos(input, op_code_pos+1, mode_first)
	var secondArgPos = calcPos(input, op_code_pos+2, mode_second)
	var targetPos = calcPos(input, op_code_pos+3, out_mode)
	var value = 0
	if input[firstArgPos] < input[secondArgPos] {
		value = 1
	}
	return targetPos, value
}

func calcMul(input []int, op_code_pos, mode_first, mode_second, out_mode int) (int, int) {
	var firstArgPos = calcPos(input, op_code_pos+1, mode_first)
	var secondArgPos = calcPos(input, op_code_pos+2, mode_second)
	var targetPos = calcPos(input, op_code_pos+3, out_mode)
	var value = input[firstArgPos] * input[secondArgPos]
	return targetPos, value
}

func calcSum(input []int, op_code_pos, mode_first, mode_second, out_mode int) (int, int) {
	var firstArgPos = calcPos(input, op_code_pos+1, mode_first)
	var secondArgPos = calcPos(input, op_code_pos+2, mode_second)
	var targetPos = calcPos(input, op_code_pos+3, out_mode)
	var value = input[firstArgPos] + input[secondArgPos]
	return targetPos, value
}

