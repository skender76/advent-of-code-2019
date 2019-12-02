package main

import (
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func calculate( noun int, verb int) int {
	var input = [...]int{1,12,2,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,6,19,23,2,23,6,27,1,5,27,31,1,10,31,35,2,6,35,39,1,39,13,43,1,43,9,47,2,47,10,51,1,5,51,55,1,55,10,59,2,59,6,63,2,6,63,67,1,5,67,71,2,9,71,75,1,75,6,79,1,6,79,83,2,83,9,87,2,87,13,91,1,10,91,95,1,95,13,99,2,13,99,103,1,103,10,107,2,107,10,111,1,111,9,115,1,115,2,119,1,9,119,0,99,2,0,14,0}
	input[1] = noun
	input[2] = verb

	for op_code:= 0; op_code < len(input); op_code +=4{

		if input[op_code] == 1 {
			var firstArgPos = input[op_code+1]
			var secondArgPos = input[op_code+2]
			var targetPos = input[op_code+3]
			var value = input[firstArgPos] + input[secondArgPos]
			input[targetPos] = value
		} else if input[op_code] == 2 {
			var firstArgPos = input[op_code+1]
			var secondArgPos = input[op_code+2]
			var targetPos = input[op_code+3]
			var value = input[firstArgPos] * input[secondArgPos]
			input[targetPos] = value
		} else if input[op_code] == 99 {
			break
		} else {
			fmt.Println("Wrong...")
		}
	}

	return input[0]
}

func main() {

	for noun:= 0; noun < 99; noun++{
		for verb:= 0; verb < 99; verb++{
			var calculateValue = calculate(noun, verb)

			if calculateValue == 19690720 {
				var res = (100 * noun) + verb
				fmt.Println(res)
				break
			}
		}
	}
}