package main



type Asteroid struct {
	x int
	y int
	observable_asteroids int
	visibileAstroids []Asteroid
	angular_coeff_from_base float64
}

func NewAsteroid() *Asteroid {
	p := new(Asteroid)
	p.x = 0
	p.y = 0
	p.observable_asteroids = 0
	p.visibileAstroids = []Asteroid{}
	p.angular_coeff_from_base = 0.0
	return p
}

func NewAsteroidInPosition(x,y int) *Asteroid {
	p := new(Asteroid)
	p.x = x
	p.y = y
	p.observable_asteroids = 0
	p.visibileAstroids = []Asteroid{}
	p.angular_coeff_from_base = 0.0
	return p
}

func (r *Asteroid)equal(asteroid Asteroid) bool {
	isEqual := false

	if r.x == asteroid.x && r.y == asteroid.y {
		isEqual = true
	}

	return isEqual
}

func findAsteroids(asteroid_field []string) []Asteroid {
	asteroids_position := []Asteroid{}
	y := 0

	for y < len(asteroid_field) {
		x := 0

		for x < len(asteroid_field[y]) {
			letter := asteroid_field[y][x]

			if letter == '#' {
				detected_asteroid := NewAsteroidInPosition(x,y)
				asteroids_position = append(asteroids_position, *detected_asteroid)
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
		base := &asteroid_positions[base_index]
		checkAsteroidVisibility(asteroid_positions, base)

		base_index++
	}
}

func checkAsteroidVisibility(asteroid_positions []Asteroid, base *Asteroid) {
	asteroid_index := 0
	for asteroid_index < len(asteroid_positions) {
		asteroid := asteroid_positions[asteroid_index]
		if !asteroid.equal(*base) {
			isSeen := true
			obstacle_index := 0
			for obstacle_index < len(asteroid_positions) {
				obstacle := asteroid_positions[obstacle_index]
				if !obstacle.equal(*base) && !obstacle.equal(asteroid) {
					isSeen = isSeen && canSee(*base, obstacle, asteroid)
				}
				obstacle_index++
			}
			if isSeen {
				base.observable_asteroids++
				base.visibileAstroids = append(base.visibileAstroids, asteroid)
			}
		}
		asteroid_index++
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
