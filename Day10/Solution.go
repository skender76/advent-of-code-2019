package main

import "fmt"

func main() {
	//var asteroid_map = []string{
	//	".#........",
	//	"..#.......",
	//	"...#......"}
	//
	//n_asteroids := findAsteroids(asteroid_map)
	//
	//count_observable_asteroids(n_asteroids)
	//
	//found, _ := find_base(n_asteroids)
	//
	//fmt.Println(found)
	//fmt.Println(n_asteroids)

	base := Asteroid{2,2,0}
	obstacle := Asteroid{ 1, 1, 0}
	target := Asteroid {3,3,0}

	isSeen := canSee(base, obstacle, target)

	fmt.Println(isSeen)

}
