package main

import (
	"fmt"
	"strings"
)

type Planet struct {
	Name           string
	orbitalPlanets []Planet
}

//func check_key_in_map(  dict map[string]Planet, key string ) (bool, Planet) {
//	var foundPlanet = Planet{"",nil}
//	var foundKey = false
//	if val, ok := dict[key]; ok {
//		foundPlanet = val
//		foundKey = true
//	}
//	return (foundKey, foundPlanet)
//}

var HEAD = Planet{"COM", nil}

func calculateOrbit(planet *Planet, orbit_map []string) {
	for _, orbit := range orbit_map {
		var orbit_pair = strings.Split(orbit,")")

		if planet.Name == orbit_pair[0] {
			var orbitalPlanets = Planet{orbit_pair[1], nil}
			planet.orbitalPlanets = append(planet.orbitalPlanets, orbitalPlanets)
		}
	}

	for index, _ := range planet.orbitalPlanets {
		calculateOrbit(&planet.orbitalPlanets[index], orbit_map)
	}
}

func count_orbit(planet Planet, steps int) int {

	var next_steps = 0

	steps++

	for index, _ := range planet.orbitalPlanets {
		next_steps += count_orbit(planet.orbitalPlanets[index], steps)
	}

	return steps + next_steps
}

func main() {
	var orbit_map = read_file("input")
	calculateOrbit(&HEAD, orbit_map)

	var result = -1
	result = count_orbit(HEAD, result)

	fmt.Println(result)
}