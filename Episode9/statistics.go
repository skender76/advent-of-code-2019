package main

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

func calcPermutationWithoutRepeatingNumber(sequence []int) [][]int {

	var combination [][]int

	for index:=0; index<5; index++ {

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

