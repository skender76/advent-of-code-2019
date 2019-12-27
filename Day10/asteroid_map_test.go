package main

import (
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

	if base.x != 3 && base.y != 4 {
		t.Errorf("Value (%d,%d) but expected (%d,%d)", base.x, base.y , 3,4)
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

	if base.x != 5 && base.y != 8 {
		t.Errorf("Value (%d,%d) but expected (%d,%d)", base.x, base.y , 5,8)
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

	if base.x != 51&& base.y != 2 {
		t.Errorf("Value (%d,%d) but expected (%d,%d)", base.x, base.y , 1,2)
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

	if base.x != 6 && base.y != 3 {
		t.Errorf("Value (%d,%d) but expected (%d,%d)", base.x, base.y , 6,3)
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

	if base.x != 11 && base.y != 13 {
		t.Errorf("Value (%d,%d) but expected (%d,%d)", base.x, base.y , 11,13)
	}

	if base.observable_asteroids != 210 {
		t.Errorf("Value %d but expected %d", base.observable_asteroids,210)
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

	if base.observable_asteroids != 282 {
		t.Errorf("Value %d but expected %d", base.observable_asteroids,282)
	}
}




