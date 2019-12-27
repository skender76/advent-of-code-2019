package main

import (
	"fmt"
	"sort"
)

func main() {
	var asteroid_map = []string{
		"###..#.##.####.##..###.#.#..",
		"#..#..###..#.......####.....",
		"#.###.#.##..###.##..#.###.#.",
		"..#.##..##...#.#.###.##.####",
		".#.##..####...####.###.##...",
		"##...###.#.##.##..###..#..#.",
		".##..###...#....###.....##.#",
		"#..##...#..#.##..####.....#.",
		".#..#.######.#..#..####....#",
		"#.##.##......#..#..####.##..",
		"##...#....#.#.##.#..#...##.#",
		"##.####.###...#.##........##",
		"......##.....#.###.##.#.#..#",
		".###..#####.#..#...#...#.###",
		"..##.###..##.#.##.#.##......",
		"......##.#.#....#..##.#.####",
		"...##..#.#.#.....##.###...##",
		".#.#..#.#....##..##.#..#.#..",
		"...#..###..##.####.#...#..##",
		"#.#......#.#..##..#...#.#..#",
		"..#.##.#......#.##...#..#.##",
		"#.##..#....#...#.##..#..#..#",
		"#..#.#.#.##..#..#.#.#...##..",
		".#...#.........#..#....#.#.#",
		"..####.#..#..##.####.#.##.##",
		".#.######......##..#.#.##.#.",
		".#....####....###.#.#.#.####",
		"....####...##.#.#...#..#.##."}

	n_asteroids := findAsteroids(asteroid_map)
	target_number_of_asteroids := 200
	base := NewAsteroidInPosition(22,19)
	hit_list := []VisibleAsteroidFromBase{}

	for len(hit_list) <= target_number_of_asteroids {
		count_observable_asteroids_from_base(n_asteroids, base)

		n_asteroids = remove_asteroid_hit(n_asteroids,base.visibileAstroids)

		sort.SliceStable(base.visibileAstroids,
			func(i, j int) bool {
				return base.visibileAstroids[i].angular_position < base.visibileAstroids[j].angular_position
			} )

		hit_list = append(hit_list, base.visibileAstroids...)

		base.visibileAstroids = []VisibleAsteroidFromBase{}
	}

	lastPos := target_number_of_asteroids - 1

	value := (hit_list[lastPos].position.x * 100) + hit_list[lastPos].position.y

	fmt.Println(value)
}
