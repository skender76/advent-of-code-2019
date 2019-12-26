package main

type Asteroid struct {
	x int
	y int
	observable_asteroids int
}

func findAsteroids(asteroid_field []string) []Asteroid {
	asteroids_position := []Asteroid{}
	y := 0

	for y < len(asteroid_field) {
		x := 0

		for x < len(asteroid_field[y]) {
			letter := asteroid_field[y][x]

			if letter == '#' {
				detected_asteroid := Asteroid{x,y, 0}
				asteroids_position = append(asteroids_position, detected_asteroid)
			}

			x++
		}

		y++
	}

	return asteroids_position
}

func checkIfObstacleBetweenOriginAndTarget(base Asteroid, obstacle Asteroid, target Asteroid) bool {
	isBetweenX := false
	isBeetweenY := false

	if base.x <= target.x {
		if base.x <= obstacle.x && obstacle.x <= target.x {
			isBetweenX = true
		}
	} else {
		if target.x < obstacle.x && obstacle.x <= base.x {
			isBetweenX = true
		}
	}

	if base.y <= target.y {
		if base.y <= obstacle.y && obstacle.y <= target.y {
			isBeetweenY = true
		}
	} else {
		if target.y < obstacle.y && obstacle.y <= base.y {
			isBeetweenY = true
		}
	}

	return isBetweenX && isBeetweenY

}

func canSee(base Asteroid, obstacle Asteroid, target Asteroid) bool {

	isSeen := true

	if base.x != obstacle.x && base.y != obstacle.y {

		if checkIfObstacleBetweenOriginAndTarget(base, obstacle, target) {
			value := ((float64((obstacle.y - base.y) * (target.x - base.x)) / float64(obstacle.x - base.x)) + float64(base.y))
			isSeen = !equal(float64(target.y),value)
		}
	} else if base.y == obstacle.y && base.y == target.y {
			if base.x < obstacle.x && target.x > obstacle.x {
				isSeen = false
			} else if base.x > obstacle.x && target.x < obstacle.x {
				isSeen = false
			}
	} else if base.x == obstacle.x && base.x == target.x {
		if base.y < obstacle.y && target.y > obstacle.y {
			isSeen = false
		} else if base.y > obstacle.y && target.y < obstacle.y {
			isSeen = false
		}
	}

	return isSeen
}

func count_observable_asteroids(asteroid_positions []Asteroid) {

	base_index := 0
	for base_index < len(asteroid_positions) {
		asteroid_index := 0
		for asteroid_index < len(asteroid_positions) {
			if asteroid_index != base_index {
				isSeen := true
				obstacle_index := 0
				for obstacle_index < len(asteroid_positions) {
					if obstacle_index != base_index && obstacle_index != asteroid_index {
						isSeen = isSeen && canSee(asteroid_positions[base_index], asteroid_positions[obstacle_index],asteroid_positions[asteroid_index])
					}
					obstacle_index++
				}
				if isSeen {
					asteroid_positions[base_index].observable_asteroids++
				}
			}
			asteroid_index++
		}

		base_index++
	}
}

func find_base(asteroid_positions []Asteroid ) (bool,Asteroid) {
	max := 0
	base := Asteroid{}
	found := false
	index := 0
	for index < len(asteroid_positions) {

		if max < asteroid_positions[index].observable_asteroids {
			max = asteroid_positions[index].observable_asteroids
			found = true
			base = asteroid_positions[index]
		}

		index++
	}

	return found,base
}
