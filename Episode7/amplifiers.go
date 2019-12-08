package main

import "fmt"

func calcPermutationWithRepeatingNumber() [][]int {

	var combination [][]int

	for index:=0; index<=4; index++ {
		for index1 :=0; index1 <=4; index1++ {
			for index2 :=0; index2 <=4; index2++ {
				for index3:=0; index3<=4; index3++ {
					for index4:=0; index4<=4; index4++ {
						var curr_comb []int
						curr_comb= append(curr_comb,index)
						curr_comb= append(curr_comb,index1)
						curr_comb= append(curr_comb,index2)
						curr_comb= append(curr_comb,index3)
						curr_comb= append(curr_comb,index4)

						combination = append(combination, curr_comb)
					}
				}
			}
		}
	}

	return combination
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func clone(input []int) []int {
	var result = make([]int, len(input))
	copy(result, input)
	return result
}

func calcPermutationWithoutRepeatingNumber() [][]int {

	var combination [][]int

	for index:=0; index<5; index++ {

		var sequence = []int {0,1,2,3,4}
		var value = sequence[index]

		var tmp_sequence = clone(sequence)
		tmp_sequence = 	remove(tmp_sequence, index)

		for index1 :=0; index1 <4; index1++ {
			var value1 = tmp_sequence[index1]
			var tmp_sequence2 = clone(tmp_sequence)
			tmp_sequence2 = remove(tmp_sequence2, index1)

			for index2 :=0; index2 <3; index2++ {
				var value2 = tmp_sequence2[index2]
				var tmp_sequence3 = clone(tmp_sequence2)
				tmp_sequence3 = remove(tmp_sequence3, index2)

				for index3:=0; index3<2; index3++ {

					var value3 = tmp_sequence3[index3]
					var tmp_sequence4 = clone(tmp_sequence3)
					tmp_sequence4 = remove(tmp_sequence4, index3)
					var value4 = tmp_sequence4[0]

					var curr_comb []int
					curr_comb= append(curr_comb,value)
					curr_comb= append(curr_comb,value1)
					curr_comb= append(curr_comb,value2)
					curr_comb= append(curr_comb,value3)
					curr_comb= append(curr_comb,value4)

					combination = append(combination, curr_comb)
				}
			}
		}
	}

	return combination
}

func calcOutputValue(settingSequence, input []int) int {

	var input_val = []int{settingSequence[0],0}
	var result = runIntcodeComputer(input_val,  input )

	input_val = []int{settingSequence[1],result}
	result = runIntcodeComputer(input_val,  input )

	input_val = []int{settingSequence[2],result}
	result = runIntcodeComputer(input_val,  input )

	input_val = []int{settingSequence[3],result}
	result = runIntcodeComputer(input_val,  input )

	input_val = []int{settingSequence[4],result}
	result = runIntcodeComputer(input_val,  input )

	return result
}

func calcMaxThrustherSignal(input []int) int {
	var phaseSettings = calcPermutationWithoutRepeatingNumber()
	var maxSignal = 0

	for index:=0;index<len(phaseSettings);index++ {

		var thrustherSignal = calcOutputValue(phaseSettings[index], input)

		if thrustherSignal > maxSignal {
			maxSignal = thrustherSignal
		}
	}

	return maxSignal
}

func main() {
	var input = []int{3,8,1001,8,10,8,105,1,0,0,21,34,59,68,89,102,183,264,345,426,99999,3,9,102,5,9,9,1001,9,5,9,4,9,99,3,9,101,3,9,9,1002,9,5,9,101,5,9,9,1002,9,3,9,1001,9,5,9,4,9,99,3,9,101,5,9,9,4,9,99,3,9,102,4,9,9,101,3,9,9,102,5,9,9,101,4,9,9,4,9,99,3,9,1002,9,5,9,1001,9,2,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,99,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,99}
	var max = calcMaxThrustherSignal(input)

	fmt.Println(max)
}
