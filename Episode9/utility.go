package main

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func clone(input []int) []int {
	var result = make([]int, len(input))
	copy(result, input)
	return result
}

func cloneExt(input []int, size int) []int {
	var result = make([]int, size)
	copy(result, input)
	return result
}
