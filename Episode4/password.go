package main

import "fmt"

func validate_pwd(pwd int) bool {

	var counter = 0
	var divident = 100000
	var double_found = false
	var incrementing_sequence = true
	var double_group_size = 1

	for counter < 5 {
		var first = pwd / divident
		pwd = pwd % divident
		divident = divident / 10
		var second = 0
		if divident == 1 {
			second = pwd
		} else {
			second = pwd / divident
		}

		if first > second {
			incrementing_sequence = false
		}

		if first == second {
			double_group_size++
		} else {
			if double_group_size == 2 {
				double_found = true
			}
			double_group_size = 1
		}

		counter++
	}

	if double_group_size == 2 {
		double_found = true
	}

	return double_found && incrementing_sequence

}

func main() {
	var i int
	var valid_pwd_counter = 0
	var result = false

	result = validate_pwd(111111)
	if result {
		fmt.Println("Valid:", 111111)
	} else {
		fmt.Println("Not valid")
	}

	result = validate_pwd(223450)
	if result {
		fmt.Println("Valid:", 223450)
	} else {
		fmt.Println("Not valid")
	}

	result = validate_pwd(123789)
	if result {
		fmt.Println("Valid:", 123789)
	} else {
		fmt.Println("Not valid")
	}

	result = validate_pwd(112233)
	if result {
		fmt.Println("Valid:", 112233)
	} else {
		fmt.Println("Not valid")
	}

	result = validate_pwd(123444)
	if result {
		fmt.Println("Valid:", 123444)
	} else {
		fmt.Println("Not valid")
	}

	result = validate_pwd(111122)
	if result {
		fmt.Println("Valid:", 111122)
	} else {
		fmt.Println("Not valid")
	}

	for i=130254; i<=678275; i++ {
		result = validate_pwd(i)
		if result {
			valid_pwd_counter++
		}
	}

	fmt.Println("Number Valid Pwd:", valid_pwd_counter)

}