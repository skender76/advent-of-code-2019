package main

import "math"

type AsteroidPosition struct {
	x int
	y int
}

func NewAsteroidPosition(x,y int) *AsteroidPosition {
	pos := new(AsteroidPosition)
	pos.x = x
	pos.y = y
	return pos
}

type VisibleAsteroidFromBase struct {
	position AsteroidPosition
	base_position AsteroidPosition
	angular_position float64
}

func calcAngle(asteroid, base AsteroidPosition) float64 {
	tan := float64(asteroid.y - base.y) / float64(asteroid.x - base.x)
	angle := math.Pi

	if base.x <= asteroid.x {
		angle = math.Pi/2 + math.Atan(tan)
	} else {
		angle = ((3*math.Pi)/2) + math.Atan(tan)
	}

	return  angle
}

func NewVisibleAsteroidFromBase(asteroid, base AsteroidPosition) *VisibleAsteroidFromBase {
	pos := new(VisibleAsteroidFromBase)
	pos.position = asteroid
	pos.base_position = base

	pos.angular_position = calcAngle(asteroid, base)

	return pos
}


type Asteroid struct {
	position AsteroidPosition
	observable_asteroids int
	visibileAstroids []VisibleAsteroidFromBase
	angular_coeff_from_base float64
}

func NewAsteroid() *Asteroid {
	p := new(Asteroid)
	p.position = *NewAsteroidPosition(0,0)
	p.observable_asteroids = 0
	p.visibileAstroids = []VisibleAsteroidFromBase{}
	p.angular_coeff_from_base = 0.0
	return p
}

func NewAsteroidInPosition(x,y int) *Asteroid {
	p := new(Asteroid)
	p.position = *NewAsteroidPosition(x,y)
	p.observable_asteroids = 0
	p.visibileAstroids = []VisibleAsteroidFromBase{}
	p.angular_coeff_from_base = 0.0
	return p
}

func (r *Asteroid)equal(asteroid Asteroid) bool {
	isEqual := false

	if r.position.x == asteroid.position.x && r.position.y == asteroid.position.y {
		isEqual = true
	}

	return isEqual
}

func (r *Asteroid)isHitBy(positionHit VisibleAsteroidFromBase) bool {
	isHit := false

	if r.position.x == positionHit.position.x && r.position.y == positionHit.position.y {
		isHit = true
	}

	return isHit
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

func remove_asteroid_hit(asteroid_field []Asteroid, asteroid_hit []VisibleAsteroidFromBase) []Asteroid {
	new_map := []Asteroid{}

	asteroid_counter := 0

	for asteroid_counter < len(asteroid_field) {
		isHit := false
		asteroid_hit_counter := 0

		for asteroid_hit_counter < len(asteroid_hit) {
			if asteroid_field[asteroid_counter].isHitBy(asteroid_hit[asteroid_hit_counter]) {
				isHit = isHit || true
			}

			asteroid_hit_counter++
		}

		if !isHit {
			new_map = append(new_map, asteroid_field[asteroid_counter])
		}

		asteroid_counter++
	}



	return new_map
}

func checkIfObstacleBetweenOriginAndTarget(base Asteroid, obstacle Asteroid, target Asteroid) bool {
	isBetweenX := false
	isBeetweenY := false

	if base.position.x <= target.position.x {
		if base.position.x <= obstacle.position.x && obstacle.position.x <= target.position.x {
			isBetweenX = true
		}
	} else {
		if target.position.x < obstacle.position.x && obstacle.position.x <= base.position.x {
			isBetweenX = true
		}
	}

	if base.position.y <= target.position.y {
		if base.position.y <= obstacle.position.y && obstacle.position.y <= target.position.y {
			isBeetweenY = true
		}
	} else {
		if target.position.y < obstacle.position.y && obstacle.position.y <= base.position.y {
			isBeetweenY = true
		}
	}

	return isBetweenX && isBeetweenY

}

func canSee(base Asteroid, obstacle Asteroid, target Asteroid) bool {

	isSeen := true

	if base.position.x != obstacle.position.x && base.position.y != obstacle.position.y {

		if checkIfObstacleBetweenOriginAndTarget(base, obstacle, target) {
			value := ((float64((obstacle.position.y - base.position.y) * (target.position.x - base.position.x)) / float64(obstacle.position.x - base.position.x)) + float64(base.position.y))
			isSeen = !equal(float64(target.position.y),value)
		}
	} else if base.position.y == obstacle.position.y && base.position.y == target.position.y {
			if base.position.x < obstacle.position.x && target.position.x > obstacle.position.x {
				isSeen = false
			} else if base.position.x > obstacle.position.x && target.position.x < obstacle.position.x {
				isSeen = false
			}
	} else if base.position.x == obstacle.position.x && base.position.x == target.position.x {
		if base.position.y < obstacle.position.y && target.position.y > obstacle.position.y {
			isSeen = false
		} else if base.position.y > obstacle.position.y && target.position.y < obstacle.position.y {
			isSeen = false
		}
	}

	return isSeen
}

func count_observable_asteroids(asteroid_positions []Asteroid) {

	base_index := 0
	for base_index < len(asteroid_positions) {
		base := &asteroid_positions[base_index]
		count_observable_asteroids_from_base(asteroid_positions, base)

		base_index++
	}
}

func count_observable_asteroids_from_base(asteroid_positions []Asteroid, base *Asteroid) {
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
				visibibleAsteroid := NewVisibleAsteroidFromBase(asteroid.position, base.position)
				base.visibileAstroids = append(base.visibileAstroids, *visibibleAsteroid)
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
