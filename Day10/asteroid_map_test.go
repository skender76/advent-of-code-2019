package main

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func TestMapCreation(t *testing.T) {
	var asteroid_map = []string{".#..#", ".....","#####","....#","...##"}

	var n_asteroids = findAsteroids(asteroid_map)
	if len(n_asteroids) != 10 {
		t.Errorf("Value %d but expected %d", len(n_asteroids) , 10) // to indicate test failed
	}
}

func TestCannotSeeTargetBehindTheObstacle(t *testing.T) {

	base := NewAsteroid()
	obstacle := NewAsteroidInPosition( 2, 0)
	target := NewAsteroidInPosition( 3, 0)

	isSeen := canSee(*base, *obstacle, *target)

	if isSeen {
		t.Errorf("Value %t but expected %t", isSeen , false) // to indicate test failed
	}
}

func TestCanSeeTargetInFrontOfTheObstacle(t *testing.T) {

	base := NewAsteroid()
	obstacle := NewAsteroidInPosition( 2, 0)
	target := NewAsteroidInPosition( 1, 0)

	isSeen := canSee(*base, *obstacle, *target)

	if !isSeen {
		t.Errorf("Value %t but expected %t", isSeen , true) // to indicate test failed
	}
}

func TestCanSeenTargetBehindOrigin(t *testing.T) {

	base := NewAsteroidInPosition( 5, 0)
	obstacle := NewAsteroidInPosition( 7, 0)
	target := NewAsteroidInPosition( 4, 0)

	isSeen := canSee(*base, *obstacle, *target)

	if !isSeen {
		t.Errorf("Value %t but expected %t", isSeen , true) // to indicate test failed
	}
}

func TestCanSeenTargetFarBehindOrigin(t *testing.T) {

	base := NewAsteroidInPosition( 5, 0)
	obstacle := NewAsteroidInPosition( 7, 0)
	target := NewAsteroidInPosition( -8, 0)

	isSeen := canSee(*base, *obstacle, *target)

	if !isSeen {
		t.Errorf("Value %t but expected %t", isSeen , true) // to indicate test failed
	}
}

func TestCannotSeeTargetBehindTheObstacleY(t *testing.T) {

	base := NewAsteroid()
	obstacle := NewAsteroidInPosition( 0, 2)
	target := NewAsteroidInPosition( 0, 3)

	isSeen := canSee(*base, *obstacle, *target)

	if isSeen {
		t.Errorf("Value %t but expected %t", isSeen , false) // to indicate test failed
	}
}

func TestCanSeeTargetInFrontOfTheObstacleY(t *testing.T) {

	base := NewAsteroid()
	obstacle := NewAsteroidInPosition( 0, 2)
	target := NewAsteroidInPosition( 0, 1)

	isSeen := canSee(*base, *obstacle, *target)

	if !isSeen {
		t.Errorf("Value %t but expected %t", isSeen , true) // to indicate test failed
	}
}

func TestCanSeenTargetBehindOriginY(t *testing.T) {

	base := NewAsteroidInPosition( 0, 5)
	obstacle := NewAsteroidInPosition( 0, 7)
	target := NewAsteroidInPosition( 0, 4)

	isSeen := canSee(*base, *obstacle, *target)

	if !isSeen {
		t.Errorf("Value %t but expected %t", isSeen , true) // to indicate test failed
	}
}

func TestCanSeenTargetFarBehindOriginY(t *testing.T) {

	base := NewAsteroidInPosition( 0, 5)
	obstacle := NewAsteroidInPosition( 0, 7)
	target := NewAsteroidInPosition( 0, -8)

	isSeen := canSee(*base, *obstacle, *target)

	if !isSeen {
		t.Errorf("Value %t but expected %t", isSeen , true) // to indicate test failed
	}
}

func TestCanSeenTargetBehindOriginXY(t *testing.T) {

	base := NewAsteroidInPosition( 1, 0)
	obstacle := NewAsteroidInPosition( 1, 2)
	target := NewAsteroidInPosition( 4, 4)

	isSeen := canSee(*base, *obstacle, *target)

	if !isSeen {
		t.Errorf("Value %t but expected %t", isSeen , true)
	}
}

func TestCannotSeenTargetBehindOriginXY(t *testing.T) {

	base := NewAsteroid()
	obstacle := NewAsteroidInPosition( 3, 2)
	target := NewAsteroidInPosition( 6, 2)

	isSeen := canSee(*base, *obstacle, *target)

	if !isSeen {
		t.Errorf("Value %t but expected %t", isSeen , true) // to indicate test failed
	}
}

func TestSimmetry(t *testing.T) {

	base := NewAsteroid()
	asteroid1 := NewAsteroidInPosition( 3, 1)
	asteroid2 := NewAsteroidInPosition( 6, 2)

	isSeen := canSee(*base, *asteroid1, *asteroid2)

	if isSeen {
		t.Errorf("Value %t but expected %t", isSeen , false) // to indicate test failed
	}

	isSeen = canSee(*base, *asteroid2, *asteroid1)

	if !isSeen {
		t.Errorf("Value %t but expected %t", isSeen , true) // to indicate test failed
	}
}

func TestSimmetryY(t *testing.T) {

	base := NewAsteroidInPosition(0, 5)
	asteroid1 := NewAsteroidInPosition( 0, 7)
	asteroid2 := NewAsteroidInPosition( 0, 8)

	isSeen := canSee(*base, *asteroid1, *asteroid2)

	if isSeen {
		t.Errorf("Value %t but expected %t", isSeen , false) // to indicate test failed
	}

	isSeen = canSee(*base, *asteroid2, *asteroid1)

	if !isSeen {
		t.Errorf("Value %t but expected %t", isSeen , true) // to indicate test failed
	}
}

func TestSimmetryX(t *testing.T) {

	base := NewAsteroidInPosition(5, 0)
	asteroid1 := NewAsteroidInPosition( 7, 0)
	asteroid2 := NewAsteroidInPosition( 8, 0)

	isSeen := canSee(*base, *asteroid1, *asteroid2)

	if isSeen {
		t.Errorf("Value %t but expected %t", isSeen , false) // to indicate test failed
	}

	isSeen = canSee(*base, *asteroid2, *asteroid1)

	if !isSeen {
		t.Errorf("Value %t but expected %t", isSeen , true) // to indicate test failed
	}
}

func TestCanSeenTargetNotAlignedOnY(t *testing.T) {

	base := NewAsteroidInPosition(5,0)
	obstacle := NewAsteroidInPosition( 7, 0)
	target := NewAsteroidInPosition( 8, 1)

	isSeen := canSee(*base, *obstacle, *target)

	if !isSeen {
		t.Errorf("Value %t but expected %t", isSeen , true) // to indicate test failed
	}
}

func TestCanSeenOnDiagonal(t *testing.T) {

	base := NewAsteroidInPosition(0,2)
	obstacle := NewAsteroidInPosition( 3, 4)
	target := NewAsteroidInPosition( 4, 4)

	isSeen := canSee(*base, *obstacle, *target)

	if !isSeen {
		t.Errorf("Value %t but expected %t", isSeen , true) // to indicate test failed
	}
}

func TestCanSeeInTheMiddleOfDiagonal(t *testing.T) {

	base := NewAsteroidInPosition(2,2)
	obstacle := NewAsteroidInPosition( 1, 1)
	target := NewAsteroidInPosition( 3, 3)

	isSeen := canSee(*base, *obstacle, *target)

	if !isSeen {
		t.Errorf("Value %t but expected %t", isSeen , true) // to indicate test failed
	}
}

func TestScenario1(t *testing.T) {
	var asteroid_map = []string{
		".#..#",
		".....",
		"#####",
		"....#",
		"...##"}

	n_asteroids := findAsteroids(asteroid_map)

	count_observable_asteroids(n_asteroids)

	found, base := find_base(n_asteroids)

	if !found {
		t.Errorf("Value %t but expected %t", found , true)
	}

	if base.position.x != 3 && base.position.y != 4 {
		t.Errorf("Value (%d,%d) but expected (%d,%d)", base.position.x, base.position.y , 3,4)
	}

	if base.observable_asteroids != 8 {
		t.Errorf("Value %d but expected %d", base.observable_asteroids,8)
	}
}

func TestScenario2(t *testing.T) {
	var asteroid_map = []string{
	"......#.#.",
	"#..#.#....",
	"..#######.",
	".#.#.###..",
	".#..#.....",
	"..#....#.#",
	"#..#....#.",
	".##.#..###",
	"##...#..#.",
	".#....####"}

	n_asteroids := findAsteroids(asteroid_map)

	count_observable_asteroids(n_asteroids)

	found, base := find_base(n_asteroids)

	if !found {
		t.Errorf("Value %t but expected %t", found , true)
	}

	if base.position.x != 5 && base.position.y != 8 {
		t.Errorf("Value (%d,%d) but expected (%d,%d)", base.position.x, base.position.y , 5,8)
	}

	if base.observable_asteroids != 33 {
		t.Errorf("Value %d but expected %d", base.observable_asteroids,33)
	}
}

func TestScenario3(t *testing.T) {
	var asteroid_map = []string{
		"#.#...#.#.",
		".###....#.",
		".#....#...",
		"##.#.#.#.#",
		"....#.#.#.",
		".##..###.#",
		"..#...##..",
		"..##....##",
		"......#...",
		".####.###."}

	n_asteroids := findAsteroids(asteroid_map)

	count_observable_asteroids(n_asteroids)

	found, base := find_base(n_asteroids)

	if !found {
		t.Errorf("Value %t but expected %t", found , true)
	}

	if base.position.x != 51&& base.position.y != 2 {
		t.Errorf("Value (%d,%d) but expected (%d,%d)", base.position.x, base.position.y , 1,2)
	}

	if base.observable_asteroids != 35 {
		t.Errorf("Value %d but expected %d", base.observable_asteroids,35)
	}
}

func TestScenario4(t *testing.T) {
	var asteroid_map = []string{
		".#..#..###",
		"####.###.#",
		"....###.#.",
		"..###.##.#",
		"##.##.#.#.",
		"....###..#",
		"..#.#..#.#",
		"#..#.#.###",
		".##...##.#",
		".....#.#.."}

	n_asteroids := findAsteroids(asteroid_map)

	count_observable_asteroids(n_asteroids)

	found, base := find_base(n_asteroids)

	if !found {
		t.Errorf("Value %t but expected %t", found , true)
	}

	if base.position.x != 6 && base.position.y != 3 {
		t.Errorf("Value (%d,%d) but expected (%d,%d)", base.position.x, base.position.y , 6,3)
	}

	if base.observable_asteroids != 41 {
		t.Errorf("Value %d but expected %d", base.observable_asteroids,41)
	}
}

func TestScenario5(t *testing.T) {
	var asteroid_map = []string{
		".#..##.###...#######",
		"##.############..##.",
		".#.######.########.#",
		".###.#######.####.#.",
		"#####.##.#.##.###.##",
		"..#####..#.#########",
		"####################",
		"#.####....###.#.#.##",
		"##.#################",
		"#####.##.###..####..",
		"..######..##.#######",
		"####.##.####...##..#",
		".#####..#.######.###",
		"##...#.##########...",
		"#.##########.#######",
		".####.#.###.###.#.##",
		"....##.##.###..#####",
		".#.#.###########.###",
		"#.#.#.#####.####.###",
		"###.##.####.##.#..##"}

	n_asteroids := findAsteroids(asteroid_map)

	count_observable_asteroids(n_asteroids)

	found, base := find_base(n_asteroids)

	if !found {
		t.Errorf("Value %t but expected %t", found , true)
	}

	if base.position.x != 11 && base.position.y != 13 {
		t.Errorf("Value (%d,%d) but expected (%d,%d)", base.position.x, base.position.y , 11,13)
	}

	if base.observable_asteroids != 210 {
		t.Errorf("Value %d but expected %d", base.observable_asteroids,210)
	}
}



func TestScenario1Quiz2(t *testing.T) {
	var asteroid_map = []string{
		".#....#####...#..",
		"##...##.#####..##",
		"##...#...#.#####.",
		"..#.....#...###..",
		"..#.#.....#....##"}

	n_asteroids := findAsteroids(asteroid_map)
	target_number_of_asteroids := 8
	base := NewAsteroidInPosition(8,4)
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

	if hit_list[lastPos].position.x != 12 && hit_list[lastPos].position.y != 2 {
		t.Errorf("Value (%d,%d) but expected (%d,%d)", hit_list[lastPos].position.x,
			hit_list[lastPos].position.y , 12,2)
	}
}

func TestAngleCalculation(t *testing.T) {

	asteroid := NewAsteroidPosition(11, 12)
	base := NewAsteroidPosition(11, 13)

	angle := calcAngle(*asteroid, *base)

	if !isZero(angle) {
		t.Errorf("Angle Expected is %f while actual %f", 0.0, angle)
	}

	asteroid = NewAsteroidPosition(11, 14)
	base = NewAsteroidPosition(11, 13)

	angle = calcAngle(*asteroid, *base)

	if compare(angle, math.Pi) != 0{
		t.Errorf("Angle Expected is %f while actual %f", math.Pi, angle)
	}

	asteroid = NewAsteroidPosition(12, 13)
	base = NewAsteroidPosition(11, 13)

	angle = calcAngle(*asteroid, *base)

	if compare(angle, (math.Pi/2)) != 0{
		t.Errorf("Angle Expected is %f while actual %f", (math.Pi/2), angle)
	}

	asteroid = NewAsteroidPosition(-12, 13)
	base = NewAsteroidPosition(11, 13)

	angle = calcAngle(*asteroid, *base)

	if compare(angle, (3*math.Pi)/2) != 0{
		t.Errorf("Angle Expected is %f while actual %f", (3*math.Pi)/2, angle)
	}

}

func TestScenario5Quiz2(t *testing.T) {
	var asteroid_map = []string{
		".#..##.###...#######",
		"##.############..##.",
		".#.######.########.#",
		".###.#######.####.#.",
		"#####.##.#.##.###.##",
		"..#####..#.#########",
		"####################",
		"#.####....###.#.#.##",
		"##.#################",
		"#####.##.###..####..",
		"..######..##.#######",
		"####.##.####...##..#",
		".#####..#.######.###",
		"##...#.##########...",
		"#.##########.#######",
		".####.#.###.###.#.##",
		"....##.##.###..#####",
		".#.#.###########.###",
		"#.#.#.#####.####.###",
		"###.##.####.##.#..##"}

	n_asteroids := findAsteroids(asteroid_map)
	target_number_of_asteroids := 200
	base := NewAsteroidInPosition(11,13)
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

	if value != 802 {
		t.Errorf("Value %d but expected %d", value, 802)
	}
}

func TestQuiz1(t *testing.T) {
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

	count_observable_asteroids(n_asteroids)

	found, base := find_base(n_asteroids)

	if !found {
		t.Errorf("Value %t but expected %t", found , true)
	}

	fmt.Println("Base (",base.position.x,",", base.position.y,")")

	if base.observable_asteroids != 282 {
		t.Errorf("Value %d but expected %d", base.observable_asteroids,282)
	}
}

func TestQuiz2(t *testing.T) {
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

	if value != 802 {
		t.Errorf("Value %d but expected %d", value, 802)
	}
}




