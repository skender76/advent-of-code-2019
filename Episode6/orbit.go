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

func count_orbit_to_planet(planet Planet, steps int, dest string) (bool, int) {
	var found = false

	if planet.Name == dest {
		found = true
	} else {
		steps++

		count_orbit_to_planet(planet, steps, dest)
		for index, _ := range planet.orbitalPlanets {
			var found_dest, step_to_dest = count_orbit_to_planet(planet.orbitalPlanets[index], steps, dest)
			if found_dest {
				steps += step_to_dest
			}
		}
	}

	return found, steps
}

func calc_planets_to_dest(root Planet, target string) []string {

	var planets []string

	if root.Name == target {
		planets = append(planets, root.Name)
	} else {
		for index, _ := range root.orbitalPlanets {
			var next_planets = calc_planets_to_dest(root.orbitalPlanets[index], target)

			if len(next_planets) != 0 {
				planets = append(planets, root.Name)
				planets = append(planets, next_planets...)
			}
		}
	}
	return planets
}

func calc_orbital_transfers(root Planet, src, dest string) int {

	var planets_to_src = calc_planets_to_dest(root, src)
	var planets_to_dest = calc_planets_to_dest(root, dest)
	var number = 0

	fmt.Println(planets_to_src)
	fmt.Println(planets_to_dest)

	if len(planets_to_src) >= len(planets_to_dest) {
		number = len(planets_to_dest)
	} else {
		number = len(planets_to_src)
	}

	var index = 0
	for index=0;index < number; index++ {
		if planets_to_dest[index] != planets_to_src[index] {
			break
		}
	}

	return ( len(planets_to_src) - index ) + ( len(planets_to_dest) - index ) - 2
}

func main() {
	var orbit_map = read_file("input")
	calculateOrbit(&HEAD, orbit_map)

	var result = calc_orbital_transfers(HEAD, "YOU", "SAN")

	fmt.Println(result)
}