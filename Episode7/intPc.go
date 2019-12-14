package main

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

func convertInsCodeToStr( instCode int ) string {
	var codeStr = ""
	var instCodeExt,_,_,_ = readOpcode(instCode)
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

func readOpcode( input int ) (int, int, int, int) {

	var op_code = input % 100
	var mode_first = (input/100) % 10
	var mode_second = (input/1000) % 10
	var mode_third = (input/10000) % 10

	return op_code, mode_first, mode_second, mode_third
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

func calcJumpIfFalsePos(input []int, op_code_pos, mode_first, mode_second int) int  {
	var firstArgPos = calcPos(input, op_code_pos+1, mode_first)
	var secondArgPos = calcPos(input, op_code_pos+2, mode_second)
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

func calcJumpIfTruePos(input []int, op_code_pos, mode_first, mode_second int) int {
	var firstArgPos = calcPos(input, op_code_pos+1, mode_first)
	var secondArgPos = calcPos(input, op_code_pos+2, mode_second)
	var new_pos = op_code_pos
	var offset = OFFSET_JUMP_IF_TRUE
	if input[firstArgPos] != 0 {
		new_pos = input[secondArgPos]
	} else {
		new_pos += offset
	}
	return new_pos
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

