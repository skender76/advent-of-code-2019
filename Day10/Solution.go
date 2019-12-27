package main

import "fmt"

func main() {
	var asteroid_map = []string{
		".#..#",
		".....",
		"#####",
		"....#",
		"...##"}

	n_asteroids := findAsteroids(asteroid_map)

	count_observable_asteroids(n_asteroids)

	found, base := find_base(n_asteroids)
	if found {
		fmt.Println("Base (",base.x,",",base.y,") = ",base.observable_asteroids,"visible asteroids")
	} else {
		fmt.Println("Error: not found any base!!")
	}

}
