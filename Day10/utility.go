package main

const EPSILON float64 = 0.0000000000000001

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

func abs(value int) int {
	abs_value := value
	if value < 0 {
		abs_value = - value
	}

	return abs_value
}

func sign(value int) int {
	sign := 0
	if value < 0 {
		sign = - 1
	} else if value > 0 {
		sign = 1
	}

	return sign
}

func equal(x, y float64 ) bool {
	isEqual := false
	diff := x - y

	 if diff > -EPSILON && diff < EPSILON {
	 	isEqual = true
	 }

	 return isEqual
}