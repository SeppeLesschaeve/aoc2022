package src

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day18() {
	content, _ := os.ReadFile("input/day18.txt")
	day18Content := string(content)
	cubeState := newCubeState()
	cubes := getCubes(day18Content)
	deltas := []Cube{{1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}, {0, 0, 1}, {0, 0, -1}}
	for _, cube := range cubes {
		cubeState.pushCube(cube, deltas)
	}
	fmt.Println(cubeState.surfaceArea)
	cubeState.cullVoids(deltas)
	fmt.Println(cubeState.surfaceArea)
}

func getCubes(content string) []Cube {
	var cubes []Cube
	cubeLines := strings.Split(content, "\n")
	for _, line := range cubeLines {
		coordinatesString := strings.Split(line, ",")
		var x, y, z int
		for i, coordinateString := range coordinatesString {
			coordinate, _ := strconv.Atoi(coordinateString)
			if i == 0 {
				x = coordinate
			} else if i == 1 {
				y = coordinate
			} else {
				z = coordinate
			}
		}
		cubes = append(cubes, Cube{x, y, z})
	}
	return cubes
}

func (cube *Cube) add(other *Cube) Cube {
	return Cube{cube.x + other.x, cube.y + other.y, cube.z + other.z}
}

func minCube(cube1, cube2 *Cube) Cube {
	return Cube{min(cube1.x, cube2.x), min(cube1.y, cube2.y), min(cube1.z, cube2.z)}
}

func maxCube(cube1, cube2 *Cube) Cube {
	return Cube{max(cube1.x, cube2.x), max(cube1.y, cube2.y), max(cube1.z, cube2.z)}
}

func newCubeState() *CubeState {
	return &CubeState{cubes: make(map[Cube]bool)}
}

func (s *CubeState) pushCube(coord Cube, deltas []Cube) {
	s.cubes[coord] = true
	s.minCube = minCube(&s.minCube, &coord)
	s.maxCube = maxCube(&s.maxCube, &coord)
	for _, delta := range deltas {
		if s.cubes[coord.add(&delta)] {
			s.surfaceArea -= 1
		} else {
			s.surfaceArea += 1
		}
	}
}

func (s *CubeState) cullVoids(deltas []Cube) {
	for x := s.minCube.x + 1; x < s.maxCube.x; x++ {
		for y := s.minCube.y + 1; y < s.maxCube.y; y++ {
			for z := s.minCube.z + 1; z < s.maxCube.z; z++ {
				coord := Cube{x, y, z}
				if !s.cubes[coord] {
					s.cullVoid(coord, deltas)
				}
			}
		}
	}
}

func (s *CubeState) cullVoid(start Cube, deltas []Cube) {
	cubeVoid := make(map[Cube]bool)
	voidSa := 0
	exterior := false
	stack := []Cube{start}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cubeVoid[curr] {
			continue
		}
		cubeVoid[curr] = true
		if curr.x == s.minCube.x || curr.y == s.minCube.y || curr.z == s.minCube.z ||
			curr.x == s.maxCube.x || curr.y == s.maxCube.y || curr.z == s.maxCube.z {
			exterior = true
		}
		for _, delta := range deltas {
			coord := curr.add(&delta)
			occupied := false
			if coord.x < s.minCube.x || coord.y < s.minCube.y || coord.z < s.minCube.z ||
				coord.x > s.maxCube.x || coord.y > s.maxCube.y || coord.z > s.maxCube.z {
				occupied = true
			}
			if cubeVoid[coord] {
				voidSa -= 1
				occupied = true
			} else {
				voidSa += 1
			}
			if s.cubes[coord] {
				occupied = true
			}
			if !occupied {
				stack = append(stack, coord)
			}
		}
	}

	if !exterior {
		s.surfaceArea -= voidSa
	}

	for coord := range cubeVoid {
		s.cubes[coord] = true
	}
}
