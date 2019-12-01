package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fuelCalculation(mass int) int {
	var fuel int
	fuel = int(mass/3) - 2

	fmt.Println(fuel)

	if fuel >= 6 {
		fuel += fuelCalculation(fuel)
		fmt.Println(fuel)
	}

	return fuel
}

func main() {
	file, err := os.Open("input")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtMasses []string

	for scanner.Scan() {
		txtMasses = append(txtMasses, scanner.Text())
	}

	file.Close()

	var fuel int
	fuel = 0
	for _, eachline := range txtMasses {
		mass, err := strconv.Atoi(eachline)
		if err == nil {
			fmt.Println(mass)
			fmt.Println(int(mass / 3))
		}

		fuel += fuelCalculation(mass)

	}



	fmt.Println(fuel)

}